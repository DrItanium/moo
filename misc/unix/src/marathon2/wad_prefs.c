/*

	 wad_prefs.c
	 Wednesday, March 22, 1995 12:08:47 AM- rdm created.

*/

#include "cseries.h"
#include <string.h>

#include "map.h"
#include "wad.h"
#include "game_errors.h"

// #include "shell.h" // for refPREFERENCES_DIALOG only
#include "wad_prefs.h"

#ifdef mpwc
	#pragma segment file_io
#endif

/* ------ local defines */
#define CURRENT_PREF_WADFILE_VERSION 0

/* ------ local globals */
struct preferences_info *prefInfo= NULL;

/* ------ local prototypes */
static void load_preferences(void);

/* ------------ Entry point */
/* Open the file, and allocate whatever internal structures are necessary in the */
/*  preferences pointer.. */
boolean w_open_preferences_file(
	char *prefName, 
	unsigned long preferences_file_type) // ostype for mac, extension for dos
{
	FileError error;
	boolean success= TRUE;

	/* allocate space for our global structure to keep track of the prefs file */
	prefInfo= (struct preferences_info *) malloc(sizeof(struct preferences_info));
	if (prefInfo)
	{
		/* Clear memory */
		memset(prefInfo, 0, sizeof(struct preferences_info));
		memcpy(prefInfo->pref_file.name, prefName, *prefName+1);

		/* check for the preferences folder using FindFolder, creating it if necessary */
		find_preferences_location(&prefInfo->pref_file);

		/* does the preferences file exist? */
		load_preferences(); /* Uses prefInfo.. */

		if(error_pending())
		{
			short type;
		
			error= get_game_error(&type);
			if(type==systemError)
			{
				switch(error)
				{
					case errFileNotFound: // to be portable!
						error= create_file(&prefInfo->pref_file, preferences_file_type);
						if (!error)
						{
							prefInfo->wad= create_empty_wad();
							set_game_error(systemError, error);
							w_write_preferences_file();
						}
						break;
						
					case errNone:
					default:
						/* Errors besides fnfErr and noErr get returned. */
						break;
				}
				
				set_game_error(systemError, error);
			} else {
				/* Something was invalid.. */
				error= delete_file(&prefInfo->pref_file);
				if(!error)
				{
					prefInfo->wad= create_empty_wad();
					set_game_error(systemError, error);
					w_write_preferences_file();
				}
				set_game_error(systemError, errNone);
			}
		}
	}
	else
	{
		set_game_error(systemError, memory_error());
	}
	
	if (error)
	{
		/* if something is broken, make sure we at least return valid prefs */
		if(prefInfo && !prefInfo->wad) prefInfo->wad= create_empty_wad();
	}

	/* Gotta bail... */
	if(!prefInfo || !prefInfo->wad)
	{
		success= FALSE;
	}
	
// dump_wad(prefInfo->wad);
	
	return success;
}

/* Caller should assert this. */
void *w_get_data_from_preferences(
	WadDataType tag,					/* Tag to read */
	short expected_size,				/* Data size */
	void (*initialize)(void *prefs),	/* Call if I have to allocate it.. */
	boolean (*validate)(void *prefs))	/* Verify function-> fixes if bad and returns TRUE */
{
	void *data;
	long length;
	
	assert(prefInfo->wad);
	
	data= extract_type_from_wad(prefInfo->wad, tag, &length);
	/* If we got the data, but it was the wrong size, or we didn't get the data... */
	if((data && length != expected_size) || (!data))
	{
		/* I have a copy of this pointer in the wad-> it's okay to do this */
		data= malloc(expected_size);
		if(data)
		{
			/* Initialize it */
			initialize(data);
			
			/* Append it to the file, for writing out.. */
			append_data_to_wad(prefInfo->wad, tag, data, expected_size, 0);
			
			/* Free our private pointer */
			free(data);
			
			/* Return the real one. */
			data= extract_type_from_wad(prefInfo->wad, tag, &length);
		}
	}
	
	if(data)
	{
		/* Only call the validation function if it is present. */
		if(validate && validate(data))
		{
			char *new_data;
			
			/* We can't hand append_data_to_wad a copy of the data pointer it */
			/* contains */
			new_data= (char *)malloc(expected_size);
			assert(new_data);
			
			memcpy(new_data, data, expected_size);
			
			/* Changed in the validation. Save to our wad. */
			append_data_to_wad(prefInfo->wad, tag, new_data, expected_size, 0);
	
			/* Free the new copy */
			free(new_data);
	
			/* And reextract the pointer.... */
			data= extract_type_from_wad(prefInfo->wad, tag, &length);
		}
	}
	
	return data;
}	

void w_write_preferences_file(
	void)
{
	struct wad_header header;
	fileref refNum;

	/* We can be called atexit. */
	if(error_pending())
	{
		set_game_error(systemError, errNone);
	}
	
	assert(!error_pending());
	refNum= open_wad_file_for_writing(&prefInfo->pref_file);
	if(!error_pending())
	{
		struct directory_entry entry;

		assert(refNum!=NONE);

		fill_default_wad_header(&prefInfo->pref_file, 
			CURRENT_WADFILE_VERSION, CURRENT_PREF_WADFILE_VERSION, 
			1, 0l, &header);
			
		if (write_wad_header(refNum, &header))
		{
			long wad_length;
			long offset= sizeof(struct wad_header);

			/* Length? */
			wad_length= calculate_wad_length(&header, prefInfo->wad);

			/* Set the entry data..... */
			set_indexed_directory_offset_and_length(&header, 
				&entry, 0, offset, wad_length, 0);
			
			/* Now write it.. */
			if (write_wad(refNum, &header, prefInfo->wad, offset))
			{
				offset+= wad_length;
				header.directory_offset= offset;
				if (write_wad_header(refNum, &header) &&
					write_directorys(refNum, &header, &entry))
				{
					/* Success! */
				} else {
					assert(error_pending());
				}
			} 
			else {
				assert(error_pending());
			}

			/* Since we don't free it, it is opened.. */
		} else {
			assert(error_pending());
		}
		close_wad_file(refNum);
	} 
	
	return;
}

static void load_preferences(
	void)
{
	fileref refNum;
	
	/* If it was already allocated, we are reloading, so free the old one.. */
	if(prefInfo->wad)
	{	
		free_wad(prefInfo->wad);
		prefInfo->wad= NULL;
	}
	
	refNum= open_wad_file_for_reading(&prefInfo->pref_file);
	if(!error_pending())
	{
		struct wad_header header;
	
		assert(refNum != NONE);
	
		/* Read the header from the wad file */
		if(read_wad_header(refNum, &header))
		{
			/* Read the indexed wad from the file */
			prefInfo->wad= read_indexed_wad_from_file(refNum, &header, 0, FALSE);
			assert(prefInfo->wad);
		}
				
		/* Close the file.. */
		close_wad_file(refNum);
	}
	
	return;
}
