/*
	PLACEMENT.C
	Friday, September 2, 1994 5:03:16 PM (ajr)
	Wednesday, July 5, 1995 8:15:57 AM- rdm cleaned up
*/

#include "cseries.h"
#include "map.h"
#include "monsters.h"
#include "items.h"

#include <string.h>

/* Constants */
#define NUMBER_OF_TICKS_BETWEEN_RECREATION (15*TICKS_PER_SECOND)
#define INVISIBLE_RANDOM_POINT_RETRIES 10

/* Global variables */
/* This is done in a single array to facilitate the saving of the game state. */
static struct object_frequency_definition object_placement_info[2*MAXIMUM_OBJECT_TYPES];
static struct object_frequency_definition *monster_placement_info;
static struct object_frequency_definition *item_placement_info;

/* private functions */
static void _recreate_objects(short object_type, short max_object_types, struct object_frequency_definition *placement_info, short *object_counts, short *random_counts);
static void add_objects(short object_class, short object_type, short count, boolean is_initial_drop);
static boolean pick_random_initial_location_of_type(short saved_type, short type, struct object_location *location);
static short pick_random_facing(short polygon_index, world_point2d *location);
static boolean choose_invisible_random_point(short *polygon_index, world_point2d *p, short object_type, boolean initial_drop);
static boolean polygon_is_valid_for_object_drop(world_point2d *location, short polygon_index, short object_type, boolean initial_drop, boolean is_random_location);

/*************************************************************************************************
 *
 * Function: load_placement_data
 * Purpose:  called by game_wad.c to get the placement information for the map.
 *
 *************************************************************************************************/
void load_placement_data(
	struct object_frequency_definition *monsters, 
	struct object_frequency_definition *items)
{
	assert(monsters != NULL && items != NULL);
	assert(NUMBER_OF_MONSTER_TYPES<=MAXIMUM_OBJECT_TYPES);
	assert(NUMBER_OF_DEFINED_ITEMS<=MAXIMUM_OBJECT_TYPES);

	item_placement_info = object_placement_info;
	monster_placement_info = object_placement_info+MAXIMUM_OBJECT_TYPES;

	/* Clear the arrays */
	memset(object_placement_info, 0, sizeof(object_placement_info));

	/* Copy them in */
	memcpy(monster_placement_info, monsters, sizeof(struct object_frequency_definition) * MAXIMUM_OBJECT_TYPES);
	memcpy(item_placement_info, items, sizeof(struct object_frequency_definition) * MAXIMUM_OBJECT_TYPES);

	memset(monster_placement_info, 0, sizeof(struct object_frequency_definition));

#ifdef DEBUG
	{
		short i;
		
		if (monster_placement_info[_monster_marine].initial_count > 0 ||
			monster_placement_info[_monster_marine].minimum_count > 0 ||
			((monster_placement_info[_monster_marine].random_count > 0 || monster_placement_info[_monster_marine].random_count == NONE) && monster_placement_info[_monster_marine].random_chance > 1))
		{
			dprintf("placement data would drop marine;g;");
		}
		
		for (i = 1; i < NUMBER_OF_MONSTER_TYPES; i++)
		{
			if (monster_placement_info[i].initial_count < 0) dprintf("bad monster initial count.;g;");
			if (monster_placement_info[i].minimum_count < 0) dprintf("bad monster minimum count.;g;");
			if (monster_placement_info[i].maximum_count < 0) dprintf("bad monster maximum count.;g;");
		}
		
		for (i = 0; i < NUMBER_OF_DEFINED_ITEMS; i++)
		{
			if (item_placement_info[i].initial_count < 0) dprintf("bad item initial count.;g;");
			if (item_placement_info[i].minimum_count < 0) dprintf("bad item minimum count.;g;");
			if (item_placement_info[i].maximum_count < 0) dprintf("bad item maximum count.;g;");
		}
	}
#endif

#if 0		
	/* Fixup network only items.. */
	if(!game_is_networked)
	{
		short index;

		/* Starts at 1 since _monster_marine is monster 0 */
		for (index= 1; index<NUMBER_OF_MONSTER_TYPES; index++)
		{
			short network_item_count;
			struct map_object *object;
			short object_index; 
			
			for(network_item_count= 0, object_index= 0, object= saved_objects; object_index<dynamic_world->initial_objects_count; ++object, ++object_index)
			{
				if(object->type==_saved_monster && object->index==index)
				{
					if(object->flags & _map_object_is_network_only)
					{
						network_item_count++;
					}
				}
			}

			monster_placement_info[index].initial_count= MAX(0, 
				monster_placement_info[index].initial_count-network_item_count)
		}

		/* Remove the network items.. */		
		for (index= 0; index<NUMBER_OF_DEFINED_ITEMS; index++)
		{
			short network_item_count;
			struct map_object *object;
			short object_index; 
			
			for (network_item_count= 0, object_index= 0, object= saved_objects;
				object_index<dynamic_world->initial_objects_count;
				++object, ++object_index)
			{
				if (object->type==_saved_item && object->index==index)
				{
					if (object->flags & _map_object_is_network_only)
					{
						network_item_count+= 1;
						object->type= NONE;
					}
				}
			}

			item_placement_info[index].initial_count= MAX(0,
				item_placement_info[index].initial_count-network_item_count);
		}
	}
#endif
	
	return;
}

/*************************************************************************************************
 *
 * Function: get_placement_info
 * Purpose:  called by game_wad.c to save the placement data.
 *
 *************************************************************************************************/
struct object_frequency_definition *get_placement_info(
	void)
{
	return object_placement_info;
}

/*************************************************************************************************
 *
 * Function: place_initial_objects
 * Purpose:  This places items and monsters according to the data that was given to me through
 *           load_placement_data().
 *
 *************************************************************************************************/
void place_initial_objects(
	void)
{
	short index;

	dynamic_world->current_civilian_causalties= dynamic_world->current_civilian_count= 0;

	if (GET_GAME_OPTIONS()&_monsters_replenish)
	{
		for (index= 1; index<NUMBER_OF_MONSTER_TYPES; index++)
		{
			if (monster_placement_info[index].initial_count)
			{
				add_objects(_object_is_monster, index, monster_placement_info[index].initial_count, TRUE);
			}
			dynamic_world->random_monsters_left[index] = monster_placement_info[index].random_count;
		}
	}
	
	for (index= 0; index<NUMBER_OF_DEFINED_ITEMS; index++)
	{
		if (item_placement_info[index].initial_count)
		{
			add_objects(_object_is_item, index, item_placement_info[index].initial_count, TRUE);
		}
		dynamic_world->random_items_left[index] = item_placement_info[index].random_count;
	}

	return;	
}

/*************************************************************************************************
 *
 * Function: mark_all_monster_collections
 * Purpose:  this needs to be called when a map is loaded to make sure the necessary collections
 *           are loaded.
 *
 *************************************************************************************************/
void mark_all_monster_collections(
	boolean loading)
{
	short index;
	struct object_frequency_definition *placement_info= monster_placement_info+1;
	
	for (index= 1; index<NUMBER_OF_MONSTER_TYPES; index++)
	{
		if (placement_info->initial_count > 0 || placement_info->minimum_count > 0 ||
			((placement_info->random_count > 0 || placement_info->random_count == NONE) && placement_info->random_chance > 1))
		{
			mark_monster_collections(index, loading);
		}
		placement_info++;
	}
	
	return;
}

void load_all_monster_sounds(
	void)
{
	short index;
	struct object_frequency_definition *placement_info= monster_placement_info+1;
	
	for (index= 1; index<NUMBER_OF_MONSTER_TYPES; index++)
	{
		if (placement_info->initial_count > 0 || placement_info->minimum_count > 0 || ((placement_info->random_count > 0 || placement_info->random_count == NONE) && placement_info->random_chance > 1))
		{
			load_monster_sounds(index);
		}
		placement_info++;
	}
	
	return;
}

/*************************************************************************************************
 *
 * Function: recreate_objects
 * Purpose:  this needs to be called periodically (probably from update_world() in marathon.c)
 *           it will periodically create new objects (monsters (not players) and items) if they
 *           need to be recreated.
 *
 *************************************************************************************************/
void recreate_objects(
	void)
{
	static long delay = 0;

	/* If time goes backwards, it means that they started a new game.  Therefore we must */
	/*  reset our delay. */
	if (dynamic_world->tick_count < delay) delay = 0;
	
	if (dynamic_world->tick_count - delay > NUMBER_OF_TICKS_BETWEEN_RECREATION)
	{
		delay= dynamic_world->tick_count;

		if (GET_GAME_OPTIONS()&_monsters_replenish)
		{
			_recreate_objects(_object_is_monster, NUMBER_OF_MONSTER_TYPES, monster_placement_info+1, 
				dynamic_world->current_monster_count, dynamic_world->random_monsters_left);
		}
		
		_recreate_objects(_object_is_item, NUMBER_OF_DEFINED_ITEMS, item_placement_info, 
			dynamic_world->current_item_count, dynamic_world->random_items_left);
	}
	
	return;
}

/*************************************************************************************************
 *
 * Function: object_was_just_added
 * Purpose:  when an object (monster or item) is created, (probably by new_monster() or new_item()
 *           this function is called. it will update my structures keep track of how many items
 *           there are.
 *
 *************************************************************************************************/
void object_was_just_added(
	short object_class, 
	short object_type)
{
	assert(object_type >= 0 && object_type < MAXIMUM_OBJECT_TYPES);
	switch(object_class)
	{
		case _object_is_monster:
			dynamic_world->current_monster_count[object_type]++;
			break;
			
		case _object_is_item:
			dynamic_world->current_item_count[object_type]++;
			break;
			
		default:
			halt(); 
			break;
	}
	
	return;
}

/*************************************************************************************************
 *
 * Function: object_was_just_destroyed
 * Purpose:  when an object (monster or item) is destroyed, this function is called. it will 
 *           update my structures keep track of how many items there are. it will also add
 *           a new item if that is necessary.
 *
 *************************************************************************************************/
void object_was_just_destroyed(
	short object_class, 
	short object_type)
{
	short diff;
	
	assert(object_type >= 0 && object_type < MAXIMUM_OBJECT_TYPES);
	
	switch(object_class)
	{
		case _object_is_monster:
			dynamic_world->current_monster_count[object_type]--;
			diff = (GET_GAME_OPTIONS()&_monsters_replenish) ? (monster_placement_info+object_type)->minimum_count - dynamic_world->current_monster_count[object_type] : 0;
			break;
			
		case _object_is_item:
			// we need to make this check because we might have destroyed an item
			// that the user was holding, but that item has a current count of 0 because
			// we never placed any on the map.
			if (dynamic_world->current_item_count[object_type]) dynamic_world->current_item_count[object_type]--;
			diff = (item_placement_info+object_type)->minimum_count - dynamic_world->current_item_count[object_type];
			break;
			
		default:
			halt();
			break;
	}
	
	if (diff>0)
	{
		add_objects(object_class, object_type, 1, FALSE);
	}
	
	return;
}

/*************************************************************************************************
 *
 * Function: get_random_player_starting_location_and_facing
 * Purpose:  returns a good place for the player to start.
 *
 *************************************************************************************************/
short get_random_player_starting_location_and_facing(
	short max_player_index,
	short team, 
	struct object_location *location)
{
	long monster_distance, player_distance;
	unsigned long best_distance;
	short starting_location_index, maximum_starting_locations, offset, index, best_index = NONE;
	struct object_location current_location;
	
	maximum_starting_locations= get_player_starting_location_and_facing(team, 0, NULL);
	offset= random() % maximum_starting_locations;
	best_distance= 0;
	
	for (starting_location_index= 0; starting_location_index<maximum_starting_locations; starting_location_index++)
	{
		index = get_player_starting_location_and_facing(team, 
			(starting_location_index+offset) % maximum_starting_locations, &current_location);

		/* Determine the distances to the nearest monster and player */
		point_is_player_visible(max_player_index, current_location.polygon_index, (world_point2d *)&current_location.p, &player_distance);
		point_is_monster_visible(current_location.polygon_index, (world_point2d *)&current_location.p, &monster_distance);
		
		if (monster_distance != 0 && player_distance != 0)
		{
			unsigned long combined_distance = player_distance + (monster_distance>>1); // weight player distance more heavily.

			if (combined_distance > best_distance)
			{
				best_index = index;
				best_distance = combined_distance;
				*location= current_location;
			}
		}
	}
	
	/* in the extremely unlikely event that there is a player or monster exactly on every location, punt */
	if (best_index == NONE)
	{
		best_index = index;
		*location= current_location;
	}

	return best_index;
}

/*------------------------------------------------------------------------------------------------
 *
 *                                   Private Functions
 *
 *------------------------------------------------------------------------------------------------/

/*************************************************************************************************
 *
 * Function: _recreate_objects
 * Purpose:  called by recreate objects, to do the actual recreation (if necessary). 
 *           This is a separate function so that it can handle both items and monsters in one loop
 *
 *************************************************************************************************/
static void _recreate_objects(
	short object_type,    // _object_is_monster or _object_is_item
	short max_object_types,  // how many monsters or objects are defined
	struct object_frequency_definition *placement_info, // from global array, probably.
	short *object_counts, // from dynamic_world
	short *random_counts) // from dynamic_world
{
	short index;
	short objects_to_add;
	boolean add_random;
	struct object_frequency_definition *indexed_placement_info= placement_info;
	
	assert(max_object_types<=MAXIMUM_OBJECT_TYPES);
	
	// it's time to check if we want to add new things.
	for (index= object_type==_object_is_monster ? 1 : 0; index < max_object_types; index++)
	{
		/* Make sure that we are at the minimum */
		objects_to_add = indexed_placement_info->minimum_count - object_counts[index];
		if (objects_to_add < 0) objects_to_add = 0;

		/* Should we add a random one? */
		if ((indexed_placement_info->random_count == NONE || random_counts[index] > 0)
			&& object_counts[index] + objects_to_add < indexed_placement_info->maximum_count
			&& random() < indexed_placement_info->random_chance)
		{
			add_random = TRUE;
			objects_to_add++;
		} else {
			add_random= FALSE;
		}

		/* If we need to add any.. */		
		if (objects_to_add)
		{
			add_objects(object_type, index, objects_to_add, FALSE);
			
			/* If we added a random, and the random_count is not NONE (which means infinite.) */
			if (add_random && indexed_placement_info->random_count != NONE) 
			{
				random_counts[index]--;
			}
		}

		indexed_placement_info++;
	}
}

/*************************************************************************************************
 *
 * Function: add_objects
 * Purpose:  This adds an object (monster or items) as many times as specified.
 *
 *************************************************************************************************/
static void add_objects(
	short object_class, 
	short object_type, 
	short count, 
	boolean is_initial_drop)
{
	short i;
	short saved_type;
	short flags;
	boolean need_random_location;
	struct object_location location;
	
	assert(object_class==_object_is_item || object_class==_object_is_monster);
	
	saved_type = (object_class == _object_is_item) ? _saved_item : _saved_monster;
	flags = (object_class == _object_is_monster) ? (monster_placement_info+object_type)->flags : (item_placement_info+object_type)->flags;
	for (i = 0; i < count; i++)
	{
		memset(&location, 0, sizeof(struct object_location));
		location.polygon_index= NONE; /* This is unnecessary, but for psychological benefits.. */
		
		need_random_location= FALSE;
		if (is_initial_drop || !(flags & _reappears_in_random_location))
		{
			if (!pick_random_initial_location_of_type(saved_type, object_type, &location))
			{
				if (is_initial_drop && (flags & _reappears_in_random_location)) need_random_location = TRUE;
				else continue;
			}
		}
		else 
		{
			need_random_location= TRUE;
		}
		
		if (need_random_location)
		{
			if (choose_invisible_random_point(&location.polygon_index, (world_point2d *)&location.p, object_class, is_initial_drop))
			{
				if (object_class == _object_is_monster) 
				{
					location.yaw= pick_random_facing(location.polygon_index, (world_point2d *)&location.p);
				}
			}
			else
			{
				continue;
			}
		}
		
		if (object_class == _object_is_item)
		{
			new_item(&location, object_type);
		}
		else
		{
			short monster_index= new_monster(&location, object_type);
			
			if (monster_index!=NONE && !is_initial_drop) 
			{
				activate_monster(monster_index);
				find_closest_appropriate_target(monster_index, TRUE);
			}
		}
	}
	
	return;
}

/*************************************************************************************************
 *
 * Function: pick_random_initial_location_of_type
 * Purpose:  this picks a place to place an object (monster or item). It picks a pre-defined
 *           starting location, but it will pick a random one.
 * Note:     this unfortunately needs _saved_item or _saved_monster instead of 
 *           _object_is_item or _object_is_monster
 *
 *************************************************************************************************/
static boolean pick_random_initial_location_of_type(
	short saved_type,
	short type,
	struct object_location *location)
{
	short              i, index, max;
	short              actual_type;
	boolean            found_location = FALSE;
	struct map_object  *saved_object;
	
	actual_type = (saved_type == _saved_item) ? _object_is_item : _object_is_monster;
	max = dynamic_world->initial_objects_count;
	index = random() % max;
	
	for (i = 0; i < max; i++)
	{
		saved_object = saved_objects + index;
		
		if (saved_object->type == saved_type && saved_object->index == type)
		{
			location->p= saved_object->location;
			location->polygon_index= saved_object->polygon_index;
			location->yaw= saved_object->facing;
			location->flags= saved_object->flags;

			if (polygon_is_valid_for_object_drop((world_point2d *)&location->p, location->polygon_index, actual_type, TRUE, FALSE))
			{
				found_location = TRUE;
				break;
			}
		}
		index = (index + 1) % max;
	} 
	
	return found_location;
}

/*************************************************************************************************
 *
 * Function: pick_random_facing
 * Purpose:  attempt to pick a good facing for an object. this is used so that monsters don't
 *           end up pointing at a wall, when pointing out into the open is a better idea.
 *
 *************************************************************************************************/
static short pick_random_facing(
	short polygon_index, 
	world_point2d *location)
{
	short          i;
	short          facing;
	short          new_polygon_index;
	world_point2d  end_point;
	
	facing= random() % NUMBER_OF_ANGLES;
	for (i= 0; i<(FULL_CIRCLE/QUARTER_CIRCLE); i++)
	{
		end_point = *location;
		translate_point2d(&end_point, WORLD_ONE, facing);
		new_polygon_index = find_new_object_polygon(location, &end_point, polygon_index);
		if (new_polygon_index != NONE)
			break;
		facing = NORMALIZE_ANGLE(facing + QUARTER_CIRCLE);	
	}
	
	return facing;
}

/*************************************************************************************************
 *
 * Function: choose_invisible_random_point
 * Purpose:  Find a place that's a good place to drop an object.
 *
 *************************************************************************************************/
static boolean choose_invisible_random_point(
	short *polygon_index, 
	world_point2d *p, 
	short object_type, 
	boolean initial_drop)
{
	short retries;
	boolean found_legal_point= FALSE;
	
	for (retries = 0; retries < INVISIBLE_RANDOM_POINT_RETRIES && !found_legal_point; ++retries)
	{
		short random_polygon_index = random() % dynamic_world->polygon_count;

		find_center_of_polygon(random_polygon_index, p);
		if(polygon_is_valid_for_object_drop(p, random_polygon_index, object_type, initial_drop, TRUE))
		{
			*polygon_index = random_polygon_index;
			found_legal_point = TRUE;
		}
	}

	return found_legal_point;
}

/*************************************************************************************************
 *
 * Function: polygon_is_valid_for_object_drop
 * Purpose:  decide if we can drop an object in this polygon.
 *
 *************************************************************************************************/
static boolean polygon_is_valid_for_object_drop(
	world_point2d *location,
	short polygon_index,
	short object_type,
	boolean initial_drop,
	boolean is_random_location)
{
	struct polygon_data *polygon = get_polygon_data(polygon_index);
	boolean valid = FALSE;
	long distance; // only used to call point_is_player_visible()

	#pragma unused (initial_drop)
	
	switch (polygon->type)
	{
		case _polygon_is_item_impassable:
			if (object_type==_object_is_item && !is_random_location)
			{
				valid= FALSE;
				break;
			}
		case _polygon_is_monster_impassable:
		case _polygon_is_platform:
		case _polygon_is_teleporter:
			if (is_random_location)
			{
				valid= FALSE;
				break;
			}
			
		default:
			if (!POLYGON_IS_DETACHED(polygon))
			{
				if (!point_is_player_visible(dynamic_world->player_count, polygon_index, location, &distance) || initial_drop)
				{
					short object_index= polygon->first_object;
					
					valid= TRUE;
					while (object_index!=NONE && valid)
					{
						struct object_data *object = get_object_data(object_index);
						
						if (is_random_location)
						{
							switch (object_type)
							{
								case _object_is_item:
									switch (GET_OBJECT_OWNER(object))
									{
										case _object_is_item:
											valid= FALSE;
									}
									break;
									
								case _object_is_monster:
									switch (GET_OBJECT_OWNER(object))
									{
										case _object_is_projectile:
										case _object_is_monster:
										case _object_is_effect:
											valid= FALSE;
									}
									break;
									
								default:
									halt();
							}
						}
						else
						{
							if (object->location.x==location->x && object->location.y==location->y)
							{
								valid= FALSE;
							}
						}
						
						object_index= object->next_object;
					}
				}
			}
	}

	return valid;
}