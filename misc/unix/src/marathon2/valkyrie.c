/*
VALKYRIE.C
Saturday, August 20, 1994 10:22:41 AM

Tuesday, August 23, 1994 4:24:21 PM
	all these VALKYRIE registers are in 32-bit address space, but that really doesn�t matter
	because the 630 is always in 32-bit mode.
Tuesday, November 29, 1994 11:27:35 AM  (Jason)
	valkyrie_erase_graphic_key_frame() doesn�t take an explicit frame or device anymore.
Tuesday, February 28, 1995 7:06:56 AM  (Jason')
	sync double-buffering to VBL; valkyrie_restore() must be called if we switch devices.
*/

/*
we need a way to determine which a) model valkyrie is installed and b) on what device
*/

#include "macintosh_cseries.h"
#include "textures.h"
#include "valkyrie.h"

#include <Retrace.h>
#include <Timer.h>

#ifdef mpwc
#pragma segment screen
#endif

/* ---------- valkyrie constants */

#define videoBuffer0 0x0
#define videoBuffer1 0x40
#define windowActive 0x20
#define pixelDouble 0x2

#define VALKYRIECLUTAddressRegister (*((volatile byte*)0x50f24000))
#define VALKYRIECLUTGraphicDataRegister (*((volatile byte*)0x50f24004))
#define VALKYRIECLUTVideoDataRegister (*((volatile byte*)0x50f24008))
#define VALKYRIECLUTColorKeyRegister (*((volatile byte*)0x50f2400c))

#define VALKYRIESubsystemConfigurationRegister (*((volatile byte*)0x50f2a00c))
#define VALKYRIEVideoInControlRegister (*((volatile byte*)0x50f2a020))
#define VALKYRIEVideoWindowXStartRegister (*((volatile word*)0x50f2a060))
#define VALKYRIEVideoWindowYStartRegister (*((volatile word*)0x50f2a064))
#define VALKYRIEVideoWindowWidthRegister (*((volatile word*)0x50f2a070))
#define VALKYRIEVideoWindowHeightRegister (*((volatile word*)0x50f2a074))
#define VALKYRIEVideoFieldStartingXPixelRegister (*((volatile word*)0x50f2a080))
#define VALKYRIEVideoFieldStartingYPixelRegister (*((volatile word*)0x50f2a084))

#define VALKYRIEVideoBufferWidth 320
#define VALKYRIEVideoBufferHeight 240
#define VALKYRIETranslatedVideoBufferRowBytes 1024

#define VALKYRIEVideoBuffer0 ((void*)0xf9097800)
#define VALKYRIEVideoBuffer1 ((void*)0xf90c0000)
#define VALKYRIETranslatedVideoBuffer0 ((void*)0xf9300000)
#define VALKYRIETranslatedVideoBuffer1 ((void*)0xf9340000)

/* ---------- structures */

struct valkyrie_data
{
	VBLTask vbl;
	
	GDHandle device;
	Rect video_frame;
	pixel8 transparent;

	short visible_video_buffer;

	VBLUPP vbl_upp;

	boolean pixel_doubling;
	boolean refresh;
};

/* ---------- globals */

volatile struct valkyrie_data *valkyrie= (struct valkyrie_data *) NULL;

/* ---------- private code */

static void valkyrie_change_clut(volatile byte *clut_register, CTabHandle color_table);

#ifdef env68k
static pascal void vbl_proc(void);
#else
static void vbl_proc(VBLTaskPtr *vbl_task);
#endif

/* ---------- code */

/* can be called multiple times */
void valkyrie_initialize(
	GDHandle device,
	boolean pixel_doubling,
	Rect *frame, // in global coordinates
	pixel8 transparent)
{
	OSErr error;

#ifdef DEBUG
	{
		Rect intersection;
		
		SectRect(frame, &(*device)->gdRect, &intersection);
		assert(EqualRect(frame, &intersection)); // must lie entirely on the given gdevice
		assert(device==GetMainDevice()); // assume valkyrie owns main device; this may not be true
	}
#endif
	
	if (!valkyrie)
	{
		valkyrie= (struct valkyrie_data *) NewPtr(sizeof(struct valkyrie_data));
		assert(valkyrie);
		
		valkyrie->vbl_upp= NewVBLProc(vbl_proc);
		assert(valkyrie->vbl_upp);

		valkyrie->vbl.qLink= (QElemPtr) NULL;
		valkyrie->vbl.qType= vType;
		valkyrie->vbl.vblAddr= valkyrie->vbl_upp;
		valkyrie->vbl.vblCount= 1;
		valkyrie->vbl.vblPhase= 0;
		
		error= SlotVInstall((QElemPtr)&valkyrie->vbl, GetSlotFromGDevice(device));
		assert(error==noErr);
	}
	
	valkyrie->video_frame= *frame;
	valkyrie->device= device;
	valkyrie->transparent= transparent;
	valkyrie->pixel_doubling= pixel_doubling;

	/* cleanup from the last guy */
	VALKYRIESubsystemConfigurationRegister= 0; /* RGB-space */
	
	valkyrie_erase_video_buffers();
	
	/* do initialization specifically relevant to us */
	VALKYRIEVideoInControlRegister= 0;
	VALKYRIEVideoFieldStartingXPixelRegister= 0;
	VALKYRIEVideoFieldStartingYPixelRegister= 0;
	VALKYRIEVideoWindowXStartRegister= frame->left;
	VALKYRIEVideoWindowYStartRegister= frame->top;
	VALKYRIEVideoWindowWidthRegister= RECTANGLE_WIDTH(frame);
	VALKYRIEVideoWindowHeightRegister= RECTANGLE_HEIGHT(frame);
	
	VALKYRIECLUTColorKeyRegister= transparent;

	/* give them the first invisible video buffer, free; we�ll only start waiting for refreshes
		after switch_to_invisible_video_buffer() is called */
	valkyrie->refresh= TRUE;

	return;
}

void valkyrie_begin(
	void)
{
	valkyrie->visible_video_buffer= 0;
	VALKYRIEVideoInControlRegister= videoBuffer0|windowActive|pixelDouble;
	
	return;
}

void valkyrie_end(
	void)
{
	VALKYRIEVideoInControlRegister= 0;
	
	return;
}

void valkyrie_restore(
	void)
{
	if (valkyrie)
	{
		OSErr error;
	
		error= SlotVRemove((QElemPtr)&valkyrie->vbl, GetSlotFromGDevice(valkyrie->device));
		assert(error==noErr);
	
		VALKYRIEVideoInControlRegister= 0;
	
		DisposePtr((Ptr)valkyrie);
		valkyrie= (struct valkyrie_data *) NULL;
	}
	
	return;
}

void valkyrie_initialize_invisible_video_buffer(
	struct bitmap_definition *bitmap)
{
	assert(valkyrie);
	
	bitmap->flags= 0;
	bitmap->width= RECTANGLE_WIDTH(&valkyrie->video_frame);
	bitmap->height= RECTANGLE_HEIGHT(&valkyrie->video_frame);
	if (valkyrie->pixel_doubling) bitmap->width>>= 1, bitmap->height>>= 1;
	bitmap->bit_depth= 16;
	bitmap->bytes_per_row= VALKYRIETranslatedVideoBufferRowBytes;
	bitmap->row_addresses[0]= valkyrie->visible_video_buffer ? (byte *)VALKYRIETranslatedVideoBuffer0 : (byte *)VALKYRIETranslatedVideoBuffer1;
	precalculate_bitmap_row_addresses(bitmap);

	while (!valkyrie->refresh);

	return;
}

void valkyrie_switch_to_invisible_video_buffer(
	void)
{
	assert(valkyrie);
	
	VALKYRIEVideoInControlRegister= windowActive | 
		(valkyrie->pixel_doubling ? pixelDouble : 0) |
		(valkyrie->visible_video_buffer ? videoBuffer0 : videoBuffer1);
	valkyrie->visible_video_buffer= !valkyrie->visible_video_buffer;
	
	/* wake up the vbl task so when know when the next refresh has ocurred (in case the
		caller asks for the invisible video buffer too soon) */
	valkyrie->refresh= FALSE;
	
	return;
}

/* if you, like, clear the screenholes the system blows chunks (mmmmm... tasty) */
void valkyrie_erase_video_buffers(
	void)
{
	short row;
	
	assert(valkyrie);
	
	for (row= 0; row<VALKYRIEVideoBufferHeight; ++row)
	{
		memset((void*)((byte*)VALKYRIETranslatedVideoBuffer0 + row*VALKYRIETranslatedVideoBufferRowBytes), 0, VALKYRIEVideoBufferWidth*sizeof(pixel16));
		memset((void*)((byte*)VALKYRIETranslatedVideoBuffer1 + row*VALKYRIETranslatedVideoBufferRowBytes), 0, VALKYRIEVideoBufferWidth*sizeof(pixel16));
	}
	
	return;
}

void valkyrie_erase_graphic_key_frame(
	pixel8 transparent_color_index)
{
	if (valkyrie)
	{
		PixMapHandle pixmap= (*valkyrie->device)->gdPMap;
		byte *base= (byte *)(*pixmap)->baseAddr;
		short bytes_per_row= (*pixmap)->rowBytes&0x3fff;
		short row;
		
		for (row= valkyrie->video_frame.top; row<valkyrie->video_frame.bottom; ++row)
		{
			memset(base + bytes_per_row*(row - (*valkyrie->device)->gdRect.top) +
				valkyrie->video_frame.left - (*valkyrie->device)->gdRect.left,
				transparent_color_index, RECTANGLE_WIDTH(&valkyrie->video_frame));
		}
	}
	
	return;
}

void valkyrie_change_video_clut(
	CTabHandle color_table)
{
	warn((*color_table)->ctSize<=PIXEL16_MAXIMUM_COMPONENT);

	if ((*color_table)->ctSize<=PIXEL16_MAXIMUM_COMPONENT)
	{
		valkyrie_change_clut(&VALKYRIECLUTVideoDataRegister, color_table);
	}
	
	return;
}

void valkyrie_change_graphic_clut(
	CTabHandle color_table)
{
	warn((*color_table)->ctSize<=PIXEL8_MAXIMUM_COLORS);

	if ((*color_table)->ctSize<=PIXEL8_MAXIMUM_COLORS)
	{
		valkyrie_change_clut(&VALKYRIECLUTGraphicDataRegister, color_table);
	}
	
	return;
}

boolean machine_has_valkyrie(
	GDSpecPtr spec)
{
	boolean has_valkyrie= FALSE;

	if (MatchGDSpec(spec)==GetMainDevice())
	{
		long machine_type;
		
		Gestalt(gestaltMachineType, &machine_type);
		switch (machine_type)
		{
			case 41:	// PPC 5200
			case 42:	// PPC 6200
			case 98:	// Q630
			case 99:	// LC580
			case 106:	// Q630 with PPC upgrade
			case 107:	// Q580 with PPC upgrade
				has_valkyrie= TRUE;
				break;
		}
	}

	return has_valkyrie;
}

/* ---------- private code */

/* ignores ColorSpec.index (assumes linear ordering) */
void valkyrie_change_clut(
	volatile byte *clut_register,
	CTabHandle color_table)
{
	short i;

	assert(valkyrie);

	VALKYRIECLUTAddressRegister= 0;
	for (i= 0; i<=(*color_table)->ctSize; ++i)
	{
		UnsignedWide ms;
		
		*clut_register= (*color_table)->ctTable[i].rgb.red>>8;
#ifdef envppc
		Microseconds(&ms);
#endif
		*clut_register= (*color_table)->ctTable[i].rgb.green>>8;
#ifdef envppc
		Microseconds(&ms);
#endif
		*clut_register= (*color_table)->ctTable[i].rgb.blue>>8;
#ifdef envppc
		Microseconds(&ms);
#endif
	}
	
	return;
}

#ifdef env68k
static pascal void vbl_proc(
	void)
{
	struct valkyrie_data *valkyrie= (struct valkyrie_data *) get_a0();
#else
static void vbl_proc(
	VBLTaskPtr *vbl_task)
{
//	struct valkyrie_data *valkyrie= (struct valkyrie_data *) vbl_task;
#endif

	valkyrie->refresh= TRUE;
	valkyrie->vbl.vblCount= 1;
	
	return;
}
