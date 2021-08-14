package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const DEBUG = false

/*
 * This is a solo paranoia game taken from the Jan/Feb issue (No 77) of
 * "SpaceGamer/FantasyGamer" magazine.
 * Article by Sam Shirley.
 * Implemented in C on Vax 11/780 under UNIX by Tim Lister
 * converted to Go by John Trammell
 * This is a public domain adventure and may not be sold for profit.
 */

const MOXIE = 13
const AGILITY = 15
const MAXKILL = 7 // The maximum number of UV's you can kill

var clone = 1
var page = 1
var computer_request = 0
var ultra_violet = false
var action_doll = false // do we have the action doll?
var hit_points = 10
var read_letter = 0
var plato_clone = 3
var blast_door = 0
var killer_count = 0

/*


 */
func get_char() rune {
	reader := bufio.NewReader(os.Stdin)
	r, _, _ := reader.ReadRune()
	return r
}

func more() {
	fmt.Printf("---------- More ----------")
	if DEBUG {
		fmt.Printf("(page %d)", page)
	}
	if get_char() == 'p' {
		character()
		fmt.Printf("---------- More ----------")
		get_char()
	}
}

func new_clone(resume int) int {
	fmt.Printf("\nClone %d just died.\n", clone)
	clone++
	if clone > 6 {
		fmt.Printf("\n*** You Lose ***\n\nAll your clones are dead.  Your name has been stricken from the records.\n\n			THE END\n")
		return 0
	} else {
		fmt.Printf("Clone %d now activated.\n", clone)
		ultra_violet = false
		action_doll = false
		hit_points = 10
		killer_count = 0
		return resume
	}
}

func dice_roll(number, faces int) int {
	total := 0
	for i := number; i > 0; i-- {
		total += rand.Int()%faces + 1
	}
	return total
}

func instructions() {
	fmt.Printf("\n\n\n\nWelcome to Paranoia!\n\n")
	fmt.Printf("HOW TO PLAY:\n\n")
	fmt.Printf("  Just press <RETURN> until you are asked to make a choice.\n")
	fmt.Printf("  Select 'a' or 'b' or whatever for your choice, then press <RETURN>.\n")
	fmt.Printf("  You may select 'p' at any time to get a display of your statistics.\n")
	fmt.Printf("  Always choose the least dangerous option.  Continue doing this until you win.\n")
	fmt.Printf("  At times you will use a skill or engage in combat and and will be informed of\n")
	fmt.Printf("  the outcome.  These sections will be self explanatory.\n\n")
	fmt.Printf("HOW TO DIE:\n\n")
	fmt.Printf("  As Philo-R-DMD you will die at times during the adventure.\n")
	fmt.Printf("  When this happens you will be given an new clone at a particular location.\n")
	fmt.Printf("  The new Philo-R will usually have to retrace some of the old Philo-R's path;\n")
	fmt.Printf("  hopefully he won't make the same mistake as his predecessor.\n\n")
	fmt.Printf("HOW TO WIN:\n\n")
	fmt.Printf("  Simply complete the mission before you expend all six clones.\n")
	fmt.Printf("  If you make it, congratulations.\n")
	fmt.Printf("  If not, you can try again later.\n")
}

func character() {
	fmt.Printf("===============================================================================\n")
	fmt.Printf("The Character : Philo-R-DMD %d\n", clone)
	fmt.Printf("Primary Attributes                      Secondary Attributes\n")
	fmt.Printf("===============================================================================\n")
	fmt.Printf("Strength ..................... 13       Carrying Capacity ................. 30\n")
	fmt.Printf("Endurance .................... 13       Damage Bonus ....................... 0\n")
	fmt.Printf("Agility ...................... 15       Macho Bonus ....................... -1\n")
	fmt.Printf("Manual Dexterity ............. 15       Melee Bonus ...................... +5%%\n")
	fmt.Printf("Moxie ........................ 13       Aimed Weapon Bonus .............. +10%%\n")
	fmt.Printf("Chutzpah ...................... 8       Comprehension Bonus .............. +4%%\n")
	fmt.Printf("Mechanical Aptitude .......... 14       Believability Bonus .............. +5%%\n")
	fmt.Printf("Power Index .................. 10       Repair Bonus ..................... +5%%\n")
	fmt.Printf("===============================================================================\n")
	fmt.Printf("Credits: 160        Secret Society: Illuminati        Secret Society Rank: 1\n")
	fmt.Printf("Service Group: Power Services               Mutant Power: Precognition\n")
	fmt.Printf("Weapon: laser pistol; to hit, 40%%; type, L; Range, 50m; Reload, 6r; Malfnt, 00\n")
	fmt.Printf("Skills: Basics 1(20%%), Aimed Weapon Combat 2(35%%), Laser 3(40%%),\n        Personal Development 1(20%%), Communications 2(29%%), Intimidation 3(34%%)\n")
	fmt.Printf("Equipment: Red Reflec Armour, Laser Pistol, Laser Barrel (red),\n")
	fmt.Printf("           Notebook & Stylus, Knife, Com Unit 1, Jump suit,\n")
	fmt.Printf("           Secret Illuminati Eye-In-The-Pyramid(tm) Decoder ring,\n")
	fmt.Printf("           Utility Belt & Pouches\n")
	fmt.Printf("===============================================================================\n")
}

func choose(a int, aptr string, b int, bptr string) int {
	fmt.Printf("\nSelect 'a' or 'b' :\n")
	fmt.Printf(" a - %s.\n b - %s.\n", aptr, bptr)
	if ch := get_char(); ch == 'a' {
		return a
	} else {
		return b
	}
}

func page1() int {
	fmt.Printf("  You wake up face down on the red and pink checked E-Z-Kleen linoleum floor.\n")
	fmt.Printf("  You recognise the pattern, it's the type preferred in the internal security\nbriefing cells.  When you finally look around you, you see that you are alone\n")
	fmt.Printf("in a large mission briefing room.\n")
	return 57
}

func page2() int {
	fmt.Printf("\"Greetings,\" says the kindly Internal Security self incrimination expert who\n")
	fmt.Printf("meets you at the door, \"How are we doing today?\"  He offers you a doughnut\n")
	fmt.Printf("and coffee and asks what brings you here.  This doesn't seem so bad, so you\n")
	fmt.Printf("tell him that you have come to confess some possible security lapses.  He\n")
	fmt.Printf("smiles knowingly, deftly catching your coffee as you slump to the floor.\n")
	fmt.Printf("\"Nothing to be alarmed about; it's just the truth serum,\" he says,\n")
	fmt.Printf("dragging you back into a discussion room.\n")
	fmt.Printf("The next five hours are a dim haze, but you can recall snatches of conversation\n")
	fmt.Printf("about your secret society, your mutant power, and your somewhat paranoid\n")
	fmt.Printf("distrust of The Computer.  This should explain why you are hogtied and moving\n")
	fmt.Printf("slowly down the conveyer belt towards the meat processing unit in Food\n")
	fmt.Printf("Services.\n")
	if computer_request == 1 {
		return new_clone(45)
	} else {
		return new_clone(32)
	}
}

func page3() int {
	fmt.Printf("You walk to the nearest Computer terminal and request more information about\n")
	fmt.Printf("Christmas.  The Computer says, \"That is an A-1 ULTRAVIOLET ONLY IMMEDIATE\n")
	fmt.Printf("TERMINATION classified topic.  What is your clearance please, Troubleshooter?\"\n")
	return choose(4, "You give your correct clearance", 5, "You lie and claim Ultraviolet clearance")
}

func page4() int {
	fmt.Printf("\"That is classified information, Troubleshooter, thank you for your inquiry.\n")
	fmt.Printf(" Please report to an Internal Security self incrimination station as soon as\n")
	fmt.Printf(" possible.\"\n")
	return 9
}

func page5() int {
	fmt.Printf("The computer says, \"Troubleshooter, you are not wearing the correct colour\n")
	fmt.Printf("uniform.  You must put on an Ultraviolet uniform immediately.  I have seen to\n")
	fmt.Printf("your needs and ordered one already; it will be here shortly.  Please wait with\n")
	fmt.Printf("your back to the wall until it arrives.\"  In less than a minute an infrared\n")
	fmt.Printf("arrives carrying a white bundle.  He asks you to sign for it, then hands it to\n")
	fmt.Printf("you and stands back, well outside of a fragmentation grenade's blast radius.\n")
	return choose(6, "You open the package and put on the uniform",
		7, "You finally come to your senses and run for it")
}

func page6() int {
	fmt.Printf("The uniform definitely makes you look snappy and pert.  It really looks\n")
	fmt.Printf("impressive, and even has the new lopsided lapel fashion that you admire so\n")
	fmt.Printf("much.  What's more, citizens of all ranks come to obsequious attention as you\n")
	fmt.Printf("walk past.  This isn't so bad being an Ultraviolet.  You could probably come\n")
	fmt.Printf("to like it, given time.\n")
	fmt.Printf("The beeping computer terminal interrupts your musings.\n")
	ultra_violet = true
	return 8
}

func page7() int {
	fmt.Printf("The corridor lights dim and are replaced by red battle lamps as the Security\n")
	fmt.Printf("Breach alarms howl all around you.  You run headlong down the corridor and\n")
	fmt.Printf("desperately windmill around a corner, only to collide with a squad of 12 Blue\n")
	fmt.Printf("clearance Vulture squadron soldiers.  \"Stop, Slime Face,\" shouts the\n")
	fmt.Printf("commander, \"or there won't be enough of you left for a tissue sample.\"\n")
	fmt.Printf("\"All right, soldiers, stuff the greasy traitor into the uniform,\" he orders,\n")
	fmt.Printf("waving the business end of his blue laser scant inches from your nose.\n")
	fmt.Printf("With his other hand he shakes open a white bundle to reveal a pristine new\n")
	fmt.Printf("Ultraviolet citizen's uniform.\n")
	fmt.Printf("One of the Vulture squadron Troubleshooters grabs you by the neck in the\n")
	fmt.Printf("exotic and very painful Vulture Clamp(tm) death grip (you saw a special about\n")
	fmt.Printf("it on the Teela O'Malley show), while the rest tear off your clothes and\n")
	fmt.Printf("force you into the Ultraviolet uniform.  The moment you are dressed they step\n")
	fmt.Printf("clear and stand at attention.\n")
	fmt.Printf("\"Thank you for your cooperation, sir,\" says the steely eyed leader of the\n")
	fmt.Printf("Vulture Squad.  \"We will be going about our business now.\"  With perfect\n")
	fmt.Printf("timing the Vultures wheel smartly and goosestep down the corridor.\n")
	fmt.Printf("Special Note: don't make the mistake of assuming that your skills have\n")
	fmt.Printf("improved any because of the uniform; you're only a Red Troubleshooter\n")
	fmt.Printf("traitorously posing as an Ultraviolet, and don't you forget it!\n")
	fmt.Printf("Suddenly, a computer terminal comes to life beside you.\n")
	ultra_violet = true
	return 8
}

func page8() int {
	fmt.Printf("\"Now, about your question, citizen.  Christmas was an old world marketing ploy\n")
	fmt.Printf("to induce lower clearance citizens to purchase vast quantities of goods, thus\n")
	fmt.Printf("accumulation a large amount of credit under the control of a single class of\n")
	fmt.Printf("citizen known as Retailers.  The strategy used is to imply that all good\n")
	fmt.Printf("citizens give gifts during Christmas, thus if one wishes to be a valuable\n")
	fmt.Printf("member of society one must also give gifts during Christmas.  More valuable\n")
	fmt.Printf("gifts make one a more valuable member, and thus did the Retailers come to\n")
	fmt.Printf("control a disproportionate amount of the currency.  In this way Christmas\n")
	fmt.Printf("eventually caused the collapse of the old world.  Understandably, Christmas\n")
	fmt.Printf("has been declared a treasonable practice in Alpha Complex.\n")
	fmt.Printf("Thank you for your inquiry.\"\n")
	fmt.Printf("You continue on your way to GDH7-beta.\n")
	return 10
}

func page9() int {
	fmt.Printf("As you walk toward the tubecar that will take you to GDH7-beta, you pass one\n")
	fmt.Printf("of the bright blue and orange Internal Security self incrimination stations.\n")
	fmt.Printf("Inside, you can see an IS agent cheerfully greet an infrared citizen and then\n")
	fmt.Printf("lead him at gunpoint into one of the rubber lined discussion rooms.\n")
	choice := choose(2, "You decide to stop here and chat, as ordered by The Computer",
		10, "You just continue blithely on past")
	if choice == 2 {
		computer_request = 1
	} else {
		computer_request = 0
	}
	return choice
}

func page10() int {
	fmt.Printf("You stroll briskly down the corridor, up a ladder, across an unrailed catwalk,\n")
	fmt.Printf("under a perilously swinging blast door in urgent need of repair, and into\n")
	fmt.Printf("tubecar grand central.  This is the bustling hub of Alpha Complex tubecar\n")
	fmt.Printf("transportation.  Before you spreads a spaghetti maze of magnalift tube tracks\n")
	fmt.Printf("and linear accelerators.  You bravely study the specially enhanced 3-D tube\n")
	fmt.Printf("route map; you wouldn't be the first Troubleshooter to take a fast tube ride\n")
	fmt.Printf("to nowhere.\n")
	if !ultra_violet {
		choice := choose(3, "You decide to ask The Computer about Christmas using a nearby terminal",
			10, "You think you have the route worked out, so you'll board a tube train")
		if choice == 3 {
			return choice
		}
	}
	fmt.Printf("You nervously select a tubecar and step aboard.\n")
	if dice_roll(2, 10) < MOXIE {
		fmt.Printf("You just caught a purple line tubecar.\n")
		return 13
	} else {
		fmt.Printf("You just caught a brown line tubecar.\n")
		return 48
	}
}

func page11() int {
	fmt.Printf("The printing on the folder says \"Experimental Self Briefing.\"\n")
	fmt.Printf("You open it and begin to read the following:\n")
	fmt.Printf("Step 1: Compel the briefing subject to attend the briefing.\n")
	fmt.Printf("        Note: See Experimental Briefing Sub Form Indigo-WY-2,\n")
	fmt.Printf("        'Experimental Self Briefing Subject Acquisition Through The Use Of\n")
	fmt.Printf("        Neurotoxin Room Foggers.'\n")
	fmt.Printf("Step 2: Inform the briefing subject that the briefing has begun.\n")
	fmt.Printf("        ATTENTION: THE BRIEFING HAS BEGUN.\n")
	fmt.Printf("Step 3: Present the briefing material to the briefing subject.\n")
	fmt.Printf("        GREETINGS TROUBLESHOOTER.\n")
	fmt.Printf("        YOU HAVE BEEN SPECIALLY SELECTED TO SINGLEHANDEDLY\n")
	fmt.Printf("        WIPE OUT A DEN OF TRAITOROUS CHRISTMAS ACTIVITY.  YOUR MISSION IS TO\n")
	fmt.Printf("        GO TO GOODS DISTRIBUTION HALL 7-BETA AND ASSESS ANY CHRISTMAS ACTIVITY\n")
	fmt.Printf("        YOU FIND THERE.  YOU ARE TO INFILTRATE THESE CHRISTMAS CELEBRANTS,\n")
	fmt.Printf("        LOCATE THEIR RINGLEADER, AN UNKNOWN MASTER RETAILER, AND BRING HIM\n")
	fmt.Printf("        BACK FOR EXECUTION AND TRIAL.  THANK YOU.  THE COMPUTER IS YOUR FRIEND.\n")
	fmt.Printf("Step 4: Sign the briefing subject's briefing release form to indicate that\n")
	fmt.Printf("        the briefing subject has completed the briefing.\n")
	fmt.Printf("        ATTENTION: PLEASE SIGN YOUR BRIEFING RELEASE FORM.\n")
	fmt.Printf("Step 5: Terminate the briefing\n")
	fmt.Printf("        ATTENTION: THE BRIEFING IS TERMINATED.\n")
	more()
	fmt.Printf("You walk to the door and hold your signed briefing release form up to the\n")
	fmt.Printf("plexiglass window.  A guard scrutinises it for a moment and then slides back\n")
	fmt.Printf("the megabolts holding the door shut.  You are now free to continue the\n")
	fmt.Printf("mission.\n")
	return choose(3, "You wish to ask The Computer for more information about Christmas",
		10, "You have decided to go directly to Goods Distribution Hall 7-beta")
}

func page12() int {
	fmt.Printf("You walk up to the door and push the button labelled \"push to exit.\"\n")
	fmt.Printf("Within seconds a surly looking guard shoves his face into the small plexiglass\n")
	fmt.Printf("window.  You can see his mouth forming words but you can't hear any of them.\n")
	fmt.Printf("You just stare at him blankly  for a few moments until he points down to a\n")
	fmt.Printf("speaker on your side of the door.  When you put your ear to it you can barely\n")
	fmt.Printf("hear him say, \"Let's see your briefing release form, bud.  You aren't\n")
	fmt.Printf("getting out of here without it.\"\n")
	return choose(11, "You sit down at the table and read the Orange packet",
		57, "You stare around the room some more")
}

func page13() int {
	fmt.Printf("You step into the shiny plasteel tubecar, wondering why the shape has always\n")
	fmt.Printf("reminded you of bullets.  The car shoots forward the instant your feet touch\n")
	fmt.Printf("the slippery gray floor, pinning you immobile against the back wall as the\n")
	fmt.Printf("tubecar careens toward GDH7-beta.  Your only solace is the knowledge that it\n")
	fmt.Printf("could be worse, much worse.\n")
	fmt.Printf("Before too long the car comes to a stop.  You can see signs for GDH7-beta\n")
	fmt.Printf("through the window.  With a little practice you discover that you can crawl\n")
	fmt.Printf("to the door and pull open the latch.\n")
	return 14
}

func page14() int {
	fmt.Printf("You manage to pull yourself out of the tubecar and look around.  Before you is\n")
	fmt.Printf("one of the most confusing things you have ever seen, a hallway that is\n")
	fmt.Printf("simultaneously both red and green clearance.  If this is the result of\n")
	fmt.Printf("Christmas then it's easy to see the evils inherent in its practice.\n")
	fmt.Printf("You are in the heart of a large goods distribution centre.  You can see all\n")
	fmt.Printf("about you evidence of traitorous secret society Christmas celebration; rubber\n")
	fmt.Printf("faced robots whiz back and forth selling toys to holiday shoppers, simul-plast\n")
	fmt.Printf("wreaths hang from every light fixture, while ahead in the shadows is a citizen\n")
	fmt.Printf("wearing a huge red synthetic flower.\n")
	return 22
}

func page15() int {
	fmt.Printf("You are set upon by a runty robot with a queer looking face and two pointy\n")
	fmt.Printf("rubber ears poking from beneath a tattered cap.  \"Hey mister,\" it says,\n")
	fmt.Printf("\"you done all your last minute Christmas shopping?  I got some real neat junk\n")
	fmt.Printf("here.  You don't wanna miss the big day tommorrow, if you know what I mean.\"\n")
	fmt.Printf("The robot opens its bag to show you a pile of shoddy Troubleshooter dolls.  It\n")
	fmt.Printf("reaches in and pulls out one of them.  \"Look, these Action Troubleshooter(tm)\n")
	fmt.Printf("dolls are the neatest thing.  This one's got moveable arms and when you\n")
	fmt.Printf("squeeze him, his little rifle squirts realistic looking napalm.  It's only\n")
	fmt.Printf("50 credits.  Oh yeah, Merry Christmas.\"\n")
	fmt.Printf("\nSelect 'a', 'b' or 'c' :\n")
	fmt.Printf(" a - You decide to buy the doll.\n")
	fmt.Printf(" b - You shoot the robot.\n")
	fmt.Printf(" c - You ignore the robot and keep searching the hall.\n")
	switch get_char() {
	case 'a':
		return 16
	case 'b':
		return 17
	case 'c':
		fallthrough
	default:
		return 22
	}
}

func page16() int {
	fmt.Printf("The doll is a good buy for fifty credits; it will make a fine Christmas present\n")
	fmt.Printf("for one of your friends.  After the sale the robot rolls away.  You can use\n")
	fmt.Printf("the doll later in combat.  It works just like a cone rifle firing napalm,\n")
	fmt.Printf("except that occasionally it will explode and blow the user to smithereens.\n")
	fmt.Printf("But don't let that stop you.\n")
	action_doll = true
	return 22
}

func page17() int {
	robot_hp := 15
	fmt.Printf("You whip out your laser and shoot the robot, but not before it squeezes the\n")
	fmt.Printf("toy at you.  The squeeze toy has the same effect as a cone rifle firing napalm,\n")
	fmt.Printf("and the elfbot's armour has no effect against your laser.\n")
	for i := 0; i < 2; i++ {
		if dice_roll(1, 100) <= 25 {
			fmt.Printf("You have been hit!\n")
			hit_points -= dice_roll(1, 10)
			if hit_points <= 0 {
				return new_clone(45)
			}
		} else {
			fmt.Printf("It missed you, but not by much!\n")
		}
		if dice_roll(1, 100) <= 40 {
			fmt.Printf("You zapped the little bastard!\n")
			robot_hp -= dice_roll(2, 10)
			if robot_hp <= 0 {
				fmt.Printf("You wasted it! Good shooting!\n")
				fmt.Printf("You will need more evidence, so you search GDH7-beta further\n")
				if hit_points < 10 {
					fmt.Printf("after the GDH medbot has patched you up.\n")
				}
				hit_points = 10
				return 22
			}
		} else {
			fmt.Printf("Damn! You missed!\n")
		}
	}
	fmt.Printf("It tried to fire again, but the toy exploded and demolished it.\n")
	fmt.Printf("You will need more evidence, so you search GDH7-beta further\n")
	if hit_points < 10 {
		fmt.Printf("after the GDH medbot has patched you up.\n")
	}
	hit_points = 10
	return 22
}

func page18() int {
	fmt.Printf("You walk to the centre of the hall, ogling like an infrared fresh from the\n")
	fmt.Printf("clone vats.  Towering before you is the most unearthly thing you have ever\n")
	fmt.Printf("seen, a green multi armed mutant horror hulking 15 feet above your head.\n")
	fmt.Printf("Its skeletal body is draped with hundreds of metallic strips (probably to\n")
	fmt.Printf("negate the effects of some insidious mutant power), and the entire hideous\n")
	fmt.Printf("creature is wrapped in a thousand blinking hazard lights.  It's times like\n")
	fmt.Printf("this when you wish you'd had some training for this job.  Luckily the\n")
	fmt.Printf("creature doesn't take notice of you but stands unmoving, as though waiting for\n")
	fmt.Printf("a summons from its dark lord, the Master Retailer.\n")
	fmt.Printf("WHAM, suddenly you are struck from behind.\n")
	if dice_roll(2, 10) < AGILITY {
		return 19
	} else {
		return 20
	}
}

func page19() int {
	fmt.Printf("Quickly you regain your balance, whirl and fire your laser into the Ultraviolet\n")
	fmt.Printf("citizen behind you.  For a moment your heart leaps to your throat, then you\n")
	fmt.Printf("realise that he is indeed dead and you will be the only one filing a report on\n")
	fmt.Printf("this incident.  Besides, he was participating in this traitorous Christmas\n")
	fmt.Printf("shopping, as is evident from the rain of shoddy toys falling all around you.\n")
	fmt.Printf("Another valorous deed done in the service of The Computer!\n")
	killer_count++
	if killer_count > (MAXKILL - clone) {
		return 21
	}
	if read_letter == 1 {
		return 22
	}
	return choose(34, "You search the body, keeping an eye open for Internal Security",
		22, "You run away like the cowardly dog you are")
}

func page20() int {
	fmt.Printf("Oh no! you can't keep your balance.  You're falling, falling head first into\n")
	fmt.Printf("the Christmas beast's gaping maw.  It's a valiant struggle; you think you are\n")
	fmt.Printf("gone when its poisonous needles dig into your flesh, but with a heroic effort\n")
	fmt.Printf("you jerk a string of lights free and jam the live wires into the creature's\n")
	fmt.Printf("spine.  The Christmas beast topples to the ground and begins to burn, filling\n")
	fmt.Printf("the area with a thick acrid smoke.  It takes only a moment to compose yourself,\n")
	fmt.Printf("and then you are ready to continue your search for the Master Retailer.\n")
	return 22
}

func page21() int {
	fmt.Printf("You have been wasting the leading citizens of Alpha Complex at a prodigious\n")
	fmt.Printf("rate.  This has not gone unnoticed by the Internal Security squad at GDH7-beta.\n")
	fmt.Printf("Suddenly, a net of laser beams spear out of the gloomy corners of the hall,\n")
	fmt.Printf("chopping you into teeny, weeny bite size pieces.\n")
	return new_clone(45)
}

func page22() int {
	fmt.Printf("You are searching Goods Distribution Hall 7-beta.\n")
	choose := dice_roll(1, 4)
	switch choose {
	case 1:
		return 18
	case 2:
		return 15
	case 3:
		return 18
	case 4:
		fallthrough
	default:
		return 29
	}
}

func page23() int {
	fmt.Printf("You go to the nearest computer terminal and declare yourself a mutant.\n")
	fmt.Printf("\"A mutant, he's a mutant,\" yells a previously unnoticed infrared who had\n")
	fmt.Printf("been looking over your shoulder.  You easily gun him down, but not before a\n")
	fmt.Printf("dozen more citizens take notice and aim their weapons at you.\n")
	return choose(28, "You tell them that it was really only a bad joke",
		24, "You want to fight it out, one against twelve")
}

func page24() int {
	fmt.Printf("Golly, I never expected someone to pick this.  I haven't even designed\n")
	fmt.Printf("the 12 citizens who are going to make a sponge out of you.  Tell you what,\n")
	fmt.Printf("I'll give you a second chance.\n")
	return choose(28, "You change your mind and say it was only a bad joke",
		25, "You REALLY want to shoot it out")
}

func page25() int {
	fmt.Printf("Boy, you really can't take a hint!\n")
	fmt.Printf("They're closing in.  Their trigger fingers are twitching, they're about to\n")
	fmt.Printf("shoot.  This is your last chance.\n")
	return choose(28, "You tell them it was all just a bad joke", 26, "You are going to shoot")
}

func page26() int {
	fmt.Printf("You can read the cold, sober hatred in their eyes (They really didn't think\n")
	fmt.Printf("it was funny), as they tighten the circle around you.  One of them shoves a\n")
	fmt.Printf("blaster up your nose, but that doesn't hurt as much as the multi-gigawatt\n")
	fmt.Printf("carbonium tipped food drill in the small of your back.\n")
	fmt.Printf("You spend the remaining micro-seconds of your life wondering what you did wrong\n")
	return new_clone(32)
}

func page27() int {
	/* doesn't exist.  Can't happen with computer version.
	designed to catch dice cheats */
	return 0
}

func page28() int {
	fmt.Printf("They don't think it's funny.\n")
	return 26
}

func page29() int {
	fmt.Printf("\"Psst, hey citizen, come here.  Pssfft,\" you hear.  When you peer around\n")
	fmt.Printf("you can see someone's dim outline in the shadows.  \"I got some information\n")
	fmt.Printf("on the Master Retailer.  It'll only cost you 30 psst credits.\"\n")
	fmt.Printf("\nSelect 'a', 'b' or 'c' :\n")
	fmt.Printf(" a - You pay the 30 credits for the info.\n")
	fmt.Printf(" b - You would rather threaten him for the information.\n")
	fmt.Printf(" c - You ignore him and walk away.\n")
	switch get_char() {
	case 'a':
		return 30
	case 'b':
		return 31
	case 'c':
		fallthrough
	default:
		return 22
	}
}

func page30() int {
	fmt.Printf("You step into the shadows and offer the man a thirty credit bill.  \"Just drop\n")
	fmt.Printf("it on the floor,\" he says.  \"So you're looking for the Master Retailer, pssfft?\n")
	fmt.Printf("I've seen him, he's a fat man in a fuzzy red and white jump suit.  They say\n")
	fmt.Printf("he's a high programmer with no respect for proper security.  If you want to\n")
	fmt.Printf("find him then pssfft step behind me and go through the door.\"\n")
	fmt.Printf("Behind the man is a reinforced plasteel blast door.  The centre of it has been\n")
	fmt.Printf("buckled toward you in a manner you only saw once before when you were field\n")
	fmt.Printf("testing the rocket assist plasma slingshot (you found it easily portable but\n")
	fmt.Printf("prone to misfire).  Luckily it isn't buckled too far for you to make out the\n")
	fmt.Printf("warning sign.  WARNING!! Don't open this door or the same thing will happen to\n")
	fmt.Printf("you.  Opening this door is a capital offense.  Do not do it.  Not at all. This\n")
	fmt.Printf("is not a joke.\n")
	fmt.Printf("\nSelect 'a', 'b' or 'c' :\n")
	fmt.Printf(" a - You use your Precognition mutant power on opening the door.\n")
	fmt.Printf(" b - You just go through the door anyway.\n")
	fmt.Printf(" c - You decide it's too dangerous and walk away.\n")
	switch get_char() {
	case 'a':
		return 56
	case 'b':
		return 33
	case 'c':
		fallthrough
	default:
		return 22
	}
}

func page31() int {
	fmt.Printf("Like any good troubleshooter you make the least expensive decision and threaten\n")
	fmt.Printf("him for information.  With lightning like reflexes you whip out your laser and\n")
	fmt.Printf("stick it up his nose.  \"Talk, you traitorous Christmas celebrator, or who nose\n")
	fmt.Printf("what will happen to you, yuk yuk,\" you pun menacingly, and then you notice\n")
	fmt.Printf("something is very wrong.  He doesn't have a nose.  As a matter of fact he's\n")
	fmt.Printf("made of one eighth inch cardboard and your laser is sticking through the other\n")
	fmt.Printf("side of his head.  \"Are you going to pay?\" says his mouth speaker,\n")
	fmt.Printf("\"or are you going to pssfft go away stupid?\"\n")
	return choose(30, "You pay the 30 credits", 22, "You pssfft go away stupid")
}

func page32() int {
	fmt.Printf("Finally it's your big chance to prove that you're as good a troubleshooter\n")
	fmt.Printf("as your previous clone.  You walk briskly to mission briefing and pick up your\n")
	fmt.Printf("previous clone's personal effects and notepad.  After reviewing the notes you\n")
	fmt.Printf("know what has to be done.  You catch the purple line to Goods Distribution Hall\n")
	fmt.Printf("7-beta and begin to search for the blast door.\n")
	return 22
}

func page33() int {
	blast_door = 1
	fmt.Printf("You release the megabolts on the blast door, then strain against it with your\n")
	fmt.Printf("awesome strength.  Slowly the door creaks open.  You bravely leap through the\n")
	fmt.Printf("opening and smack your head into the barrel of a 300 mm 'ultra shock' class\n")
	fmt.Printf("plasma cannon.  It's dark in the barrel now, but just before your head got\n")
	fmt.Printf("stuck you can remember seeing a group of technicians anxiously watch you leap\n")
	fmt.Printf("into the room.\n")
	if ultra_violet {
		return 35
	} else {
		return 36
	}
}

func page34() int {
	fmt.Printf("You have found a sealed envelope on the body.  You open it and read:\n")
	fmt.Printf("\"WARNING: Ultraviolet Clearance ONLY.  DO NOT READ.\n")
	fmt.Printf("Memo from Chico-U-MRX4 to Harpo-U-MRX5.\n")
	fmt.Printf("The planned takeover of the Troubleshooter Training Course goes well, Comrade.\n")
	fmt.Printf("Once we have trained the unwitting bourgeois troubleshooters to work as\n")
	fmt.Printf("communist dupes, the overthrow of Alpha Complex will be unstoppable.  My survey\n")
	fmt.Printf("of the complex has convinced me that no one suspects a thing; soon it will be\n")
	fmt.Printf("too late for them to oppose the revolution.  The only thing that could possibly\n")
	fmt.Printf("impede the people's revolution would be someone alerting The Computer to our\n")
	fmt.Printf("plans (for instance, some enterprising Troubleshooter could tell The Computer\n")
	fmt.Printf("that the communists have liberated the Troubleshooter Training Course and plan\n")
	fmt.Printf("to use it as a jumping off point from which to undermine the stability of all\n")
	fmt.Printf("Alpha Complex), but as we both know, the capitalistic Troubleshooters would\n")
	fmt.Printf("never serve the interests of the proletariat above their own bourgeois desires.\n")
	fmt.Printf("P.S. I'm doing some Christmas shopping later today.  Would you like me to pick\n")
	fmt.Printf("you up something?\"\n")
	more()
	fmt.Printf("When you put down the memo you are overcome by that strange deja'vu again.\n")
	fmt.Printf("You see yourself talking privately with The Computer.  You are telling it all\n")
	fmt.Printf("about the communists' plan, and then the scene shifts and you see yourself\n")
	fmt.Printf("showered with awards for foiling the insidious communist plot to take over the\n")
	fmt.Printf("complex.\n")
	read_letter = 1
	return choose(46, "You rush off to the nearest computer terminal to expose the commies", 22, "You wander off to look for more evidence")
}

func page35() int {
	fmt.Printf("\"Oh master,\" you hear through the gun barrel, \"where have you been? It is\n")
	fmt.Printf("time for the great Christmas gifting ceremony.  You had better hurry and get\n")
	fmt.Printf("the costume on or the trainee may begin to suspect.\"  For the second time\n")
	fmt.Printf("today you are forced to wear attire not of your own choosing.  They zip the\n")
	fmt.Printf("suit to your chin just as you hear gunfire erupt behind you.\n")
	fmt.Printf("\"Oh no! Who left the door open?  The commies will get in.  Quick, fire the\n")
	fmt.Printf("laser cannon or we're all doomed.\"\n")
	fmt.Printf("\"Too late you capitalist swine, the people's revolutionary strike force claims\n")
	fmt.Printf("this cannon for the proletariat's valiant struggle against oppression.  Take\n")
	fmt.Printf("that, you running dog imperialist lackey.  ZAP, KAPOW\"\n")
	fmt.Printf("Just when you think that things couldn't get worse, \"Aha, look what we have\n")
	fmt.Printf("here, the Master Retailer himself with his head caught in his own cannon.  His\n")
	fmt.Printf("death will serve as a symbol of freedom for all Alpha Complex.\n")
	fmt.Printf("Fire the cannon.\"\n")
	return new_clone(32)
}

func page36() int {
	fmt.Printf("\"Congratulations, troubleshooter, you have successfully found the lair of the\n")
	fmt.Printf("Master Retailer and completed the Troubleshooter Training Course test mission,\"\n")
	fmt.Printf("a muffled voice tells you through the barrel.  \"Once we dislodge your head\n")
	fmt.Printf("from the barrel of the 'Ultra Shock' plasma cannon you can begin with the\n")
	fmt.Printf("training seminars, the first of which will concern the 100%% accurate\n")
	fmt.Printf("identification and elimination of unregistered mutants.  If you have any\n")
	fmt.Printf("objections please voice them now.\"\n")
	fmt.Printf("\nSelect 'a', 'b' or 'c' :\n")
	fmt.Printf(" a - You appreciate his courtesy and voice an objection.\n")
	fmt.Printf(" b - After your head is removed from the cannon, you register as a mutant.\n")
	fmt.Printf(" c - After your head is removed from the cannon, you go to the unregistered\n")
	fmt.Printf("     mutant identification and elimination seminar.\n")
	switch get_char() {
	case 'a':
		return new_clone(32)
	case 'b':
		return 23
	case 'c':
		fallthrough
	default:
		return 37
	}
}

func page37() int {
	fmt.Printf("\"Come with me please, Troubleshooter,\" says the Green clearance technician\n")
	fmt.Printf("after he has dislodged your head from the cannon.  \"You have been participating\n")
	fmt.Printf("in the Troubleshooter Training Course since you got off the tube car in\n")
	fmt.Printf("GDH7-beta,\" he explains as he leads you down a corridor.  \"The entire\n")
	fmt.Printf("Christmas assignment was a test mission to assess your current level of\n")
	fmt.Printf("training.  You didn't do so well.  We're going to start at the beginning with\n")
	fmt.Printf("the other student.  Ah, here we are, the mutant identification and elimination\n")
	fmt.Printf("lecture.\"  He shows you into a vast lecture hall filled with empty seats.\n")
	fmt.Printf("There is only one other student here, a Troubleshooter near the front row\n")
	fmt.Printf("playing with his Action Troubleshooter(tm) figure.  \"Find a seat and I will\n")
	fmt.Printf("begin,\" says the instructor.\n")
	return 38
}

func page38() int {
	fmt.Printf("\"I am Plato-B-PHI%d, head of mutant propaganda here at the training course.\n", plato_clone)
	fmt.Printf("If you have any questions about mutants please come to me.  Today I will be\n")
	fmt.Printf("talking about mutant detection.  Detecting mutants is very easy.  One simply\n")
	fmt.Printf("watches for certain tell tale signs, such as the green scaly skin, the third\n")
	fmt.Printf("arm growing from the forehead, or other similar disfigurements so common with\n")
	fmt.Printf("their kind.  There are, however, a few rare specimens that show no outward sign\n")
	fmt.Printf("of their treason.  This has been a significant problem, so our researchers have\n")
	fmt.Printf("been working on a solution.  I would like a volunteer to test this device,\"\n")
	fmt.Printf("he says, holding up a ray gun looking thing.  \"It is a mutant detection ray.\n")
	fmt.Printf("This little button detects for mutants, and this big button stuns them once\n")
	fmt.Printf("they are discovered.  Who would like to volunteer for a test?\"\n")
	fmt.Printf("The Troubleshooter down the front squirms deeper into his chair.\n")
	return choose(39, "You volunteer for the test", 40, "You duck behind a chair and hope the instructor doesn't notice you")
}

func page39() int {
	fmt.Printf("You bravely volunteer to test the mutant detection gun.  You stand up and walk\n")
	fmt.Printf("down the steps to the podium, passing a very relieved Troubleshooter along the\n")
	fmt.Printf("way.  When you reach the podium Plato-B-PHI hands you the mutant detection gun\n")
	fmt.Printf("and says, \"Here, aim the gun at that Troubleshooter and push the small button.\n")
	fmt.Printf("If you see a purple light, stun him.\"  Grasping the opportunity to prove your\n")
	fmt.Printf("worth to The Computer, you fire the mutant detection ray at the Troubleshooter.\n")
	fmt.Printf("A brilliant purple nimbus instantly surrounds his body.  You slip your finger\n")
	fmt.Printf("to the large stun button and he falls writhing to the floor.\n")
	fmt.Printf("\"Good shot,\" says the instructor as you hand him the mutant detection gun,\n")
	fmt.Printf("\"I'll see that you get a commendation for this.  It seems you have the hang\n")
	fmt.Printf("of mutant detection and elimination.  You can go on to the secret society\n")
	fmt.Printf("infiltration class.  I'll see that the little mutie gets packaged for\n")
	fmt.Printf("tomorrow's mutant dissection class.\"\n")
	return 41
}

func page40() int {
	fmt.Printf("You breathe a sigh of relief as Plato-B-PHI picks on the other Troubleshooter.\n")
	fmt.Printf("\"You down here in the front,\" says the instructor pointing at the other\n")
	fmt.Printf("Troubleshooter, \"you'll make a good volunteer.  Please step forward.\"\n")
	fmt.Printf("The Troubleshooter looks around with a `who me?' expression on his face, but\n")
	fmt.Printf("since he is the only one visible in the audience he figures his number is up.\n")
	fmt.Printf("He walks down to the podium clutching his Action Troubleshooter(tm) doll before\n")
	fmt.Printf("him like a weapon.  \"Here,\" says Plato-B-PHI, \"take the mutant detection ray\n")
	fmt.Printf("and point it at the audience.  If there are any mutants out there we'll know\n")
	fmt.Printf("soon enough.\"  Suddenly your skin prickles with static electricity as a bright\n")
	fmt.Printf("purple nimbus surrounds your body.  \"Ha Ha, got one,\" says the instructor.\n")
	fmt.Printf("\"Stun him before he gets away.\"\n")
	more()
	for {
		if dice_roll(1, 100) <= 30 {
			fmt.Printf("His shot hits you.  You feel numb all over.\n")
			return 49
		} else {
			fmt.Printf("His shot just missed.\n")
		}
		if dice_roll(1, 100) <= 40 {
			fmt.Printf("You just blew his head off.  His lifeless hand drops the mutant detector ray.\n")
			return 50
		} else {
			fmt.Printf("You burnt a hole in the podium.  He sights the mutant detector ray on you.\n")
		}
	}
}

func page41() int {
	fmt.Printf("You stumble down the hallway of the Troubleshooter Training Course looking for\n")
	fmt.Printf("your next class.  Up ahead you see one of the instructors waving to you.  When\n")
	fmt.Printf("you get there he shakes your hand and says, \"I am Jung-I-PSY.  Welcome to the\n")
	fmt.Printf("secret society infiltration seminar.  I hope you ...\"  You don't catch the\n")
	fmt.Printf("rest of his greeting because you're paying too much attention to his handshake;\n")
	fmt.Printf("it is the strangest thing that has ever been done to your hand, sort of how it\n")
	fmt.Printf("would feel if you put a neuro whip in a high energy palm massage unit.\n")
	fmt.Printf("It doesn't take you long to learn what he is up to; you feel him briefly shake\n")
	fmt.Printf("your hand with the secret Illuminati handshake.\n")
	return choose(42, "You respond with the proper Illuminati code phrase, \"Ewige Blumenkraft\"", 43, "You ignore this secret society contact")
}

func page42() int {
	fmt.Printf("\"Aha, so you are a member of the elitist Illuminati secret society,\" he says\n")
	fmt.Printf("loudly, \"that is most interesting.\"  He turns to the large class already\n")
	fmt.Printf("seated in the auditorium and says, \"You see, class, by simply using the correct\n")
	fmt.Printf("hand shake you can identify the member of any secret society.  Please keep your\n")
	fmt.Printf("weapons trained on him while I call a guard.\n")
	return choose(51, "You run for it", 52, "You wait for the guard")
}

func page43() int {
	fmt.Printf("You sit through a long lecture on how to recognise and infiltrate secret\n")
	fmt.Printf("societies, with an emphasis on mimicking secret handshakes.  The basic theory,\n")
	fmt.Printf("which you realise to be sound from your Iluminati training, is that with the\n")
	fmt.Printf("proper handshake you can pass unnoticed in any secret society gathering.\n")
	fmt.Printf("What's more, the proper handshake will open doors faster than an 'ultra shock'\n")
	fmt.Printf("plasma cannon.  You are certain that with the information you learn here you\n")
	fmt.Printf("will easily be promoted to the next level of your Illuminati secret society.\n")
	fmt.Printf("The lecture continues for three hours, during which you have the opportunity\n")
	fmt.Printf("to practice many different handshakes.  Afterwards everyone is directed to\n")
	fmt.Printf("attend the graduation ceremony.  Before you must go you have a little time to\n")
	fmt.Printf("talk to The Computer about, you know, certain topics.\n")
	return choose(44, "You go looking for a computer terminal", 55, "You go to the graduation ceremony immediately")
}

func page44() int {
	fmt.Printf("You walk down to a semi-secluded part of the training course complex and\n")
	fmt.Printf("activate a computer terminal.  \"AT YOUR SERVICE\" reads the computer screen.\n")
	if read_letter == 0 {
		return choose(23, "You register yourself as a mutant",
			55, "You change your mind and go to the graduation ceremony")
	}
	fmt.Printf("\nSelect 'a', 'b' or 'c' :\n")
	fmt.Printf(" a - You register yourself as a mutant.\n")
	fmt.Printf(" b - You want to chat about the commies.\n")
	fmt.Printf(" c - You change your mind and go to the graduation ceremony.\n")
	switch get_char() {
	case 'a':
		return 23
	case 'b':
		return 46
	case 'c':
		return 55
	}
	return 55
}

func page45() int {
	fmt.Printf("\"Hrank Hrank,\" snorts the alarm in your living quarters.  Something is up.\n")
	fmt.Printf("You look at the monitor above the bathroom mirror and see the message you have\n")
	fmt.Printf("been waiting for all these years.  \"ATTENTION TROUBLESHOOTER, YOU ARE BEING\n")
	fmt.Printf("ACTIVATED. PLEASE REPORT IMMEDIATELY TO MISSION ASSIGNMENT ROOM A17/GAMMA/LB22.\n")
	fmt.Printf("THANK YOU. THE COMPUTER IS YOUR FRIEND.\"  When you arrive at mission\n")
	fmt.Printf("assignment room A17-gamma/LB22 you are given your previous clone's\n")
	fmt.Printf("remaining possessions and notebook.  You puzzle through your predecessor's\n")
	fmt.Printf("cryptic notes, managing to decipher enough to lead you to the tube station and\n")
	fmt.Printf("the tube car to GDH7-beta.\n")
	return 10
}

func page46() int {
	fmt.Printf("\"Why do you ask about the communists, Troubleshooter?  It is not in the\n")
	fmt.Printf("interest of your continued survival to be asking about such topics,\" says\n")
	fmt.Printf("The Computer.\n")
	return choose(53, "You insist on talking about the communists", 54, "You change the subject")
}

func page47() int {
	fmt.Printf("The Computer orders the entire Vulture squadron to terminate the Troubleshooter\n")
	fmt.Printf("Training Course.  Unfortunately you too are terminated for possessing\n")
	fmt.Printf("classified information.\n\n")
	fmt.Printf("Don't act so innocent, we both know that you are an Illuminatus which is in\n")
	fmt.Printf("itself an act of treason.\n\n")
	fmt.Printf("Don't look to me for sympathy.\n\n")
	fmt.Printf("			THE END\n")
	return 0
}

func page48() int {
	fmt.Printf("The tubecar shoots forward as you enter, slamming you back into a pile of\n")
	fmt.Printf("garbage.  The front end rotates upward and you, the garbage and the garbage\n")
	fmt.Printf("disposal car shoot straight up out of Alpha Complex.  One of the last things\n")
	fmt.Printf("you see is a small blue sphere slowly dwindling behind you.  After you fail to\n")
	fmt.Printf("report in, you will be assumed dead.\n")
	return new_clone(45)
}

func page49() int {
	fmt.Printf("The instructor drags your inert body into a specimen detainment cage.\n")
	fmt.Printf("\"He'll make a good subject for tomorrow's mutant dissection class,\" you hear.\n")
	return new_clone(32)
}

func page50() int {
	fmt.Printf("You put down the other Troubleshooter, and then wisely decide to drill a few\n")
	fmt.Printf("holes in the instructor as well; the only good witness is a dead witness.\n")
	fmt.Printf("You continue with the training course.\n")
	plato_clone++
	return 41
}

func page51() int {
	fmt.Printf("You run for it, but you don't run far.  Three hundred strange and exotic\n")
	fmt.Printf("weapons turn you into a freeze dried cloud of soot.\n")
	return new_clone(32)
}

func page52() int {
	fmt.Printf("You wisely wait until the instructor returns with a Blue Internal Security\n")
	fmt.Printf("guard.  The guard leads you to an Internal Security self incrimination station.\n")
	return 2
}

func page53() int {
	fmt.Printf("You tell The Computer about:\n")
	return choose(47, "The commies who have infiltrated the Troubleshooter Training Course\n     and the impending People's Revolution", 54, "Something less dangerous")
}

func page54() int {
	fmt.Printf("\"Do not try to change the subject, Troubleshooter,\" says The Computer.\n")
	fmt.Printf("\"It is a serious crime to ask about the communists.  You will be terminated\n")
	fmt.Printf("immediately.  Thank you for your inquiry.  The Computer is your friend.\"\n")
	fmt.Printf("Steel bars drop to your left and right, trapping you here in the hallway.\n")
	fmt.Printf("A spotlight beams from the computer console to brilliantly iiluminate you while\n")
	fmt.Printf("the speaker above your head rapidly repeats \"Traitor, Traitor, Traitor.\"\n")
	fmt.Printf("It doesn't take long for a few guards to notice your predicament and come to\n")
	fmt.Printf("finish you off.\n")
	if blast_door == 0 {
		return new_clone(45)
	} else {
		return new_clone(32)
	}
}

func page55() int {
	fmt.Printf("You and 300 other excited graduates are marched  from the lecture hall and into\n")
	fmt.Printf("a large auditorium for the graduation exercise.  The auditorium is\n")
	fmt.Printf("extravagantly decorated in the colours of the graduating class.  Great red and\n")
	fmt.Printf("green plasti-paper ribbons drape from the walls, while a huge sign reading\n")
	fmt.Printf("\"Congratulations class of GDH7-beta-203.44/A\" hangs from the raised stage down\n")
	fmt.Printf("front.  Once everyone finds a seat the ceremony begins.  Jung-I-PSY is the\n")
	fmt.Printf("first to speak, \"Congratulations students, you have successfully survived the\n")
	fmt.Printf("Troubleshooter Training Course.  It always brings me great pride to address\n")
	fmt.Printf("the graduating class, for I know, as I am sure you do too, that you are now\n")
	fmt.Printf("qualified for the most perilous missions The Computer may select for you.  The\n")
	fmt.Printf("thanks is not owed to us of the teaching staff, but to all of you, who have\n")
	fmt.Printf("persevered and graduated.  Good luck and die trying.\"  Then the instructor\n")
	fmt.Printf("begins reading the names of the students who one by one walk to the front of\n")
	fmt.Printf("the auditorium and receive their diplomas.  Soon it is your turn,\n")
	fmt.Printf("\"Philo-R-DMD, graduating a master of mutant identification and secret society\n")
	fmt.Printf("infiltration.\"  You walk up and receive your diploma from Plato-B-PHI%d, then\n", plato_clone)
	fmt.Printf("return to your seat.  There is another speech after the diplomas are handed\n")
	fmt.Printf("out, but it is cut short by by rapid fire laser bursts from the high spirited\n")
	fmt.Printf("graduating class.  You are free to return to your barracks to wait, trained\n")
	fmt.Printf("and fully qualified, for your next mission.  You also get that cherished\n")
	fmt.Printf("promotion from the Illuminati secret society.  In a week you receive a\n")
	fmt.Printf("detailed Training Course bill totalling 1,523 credits.\n")
	fmt.Printf("			THE END\n")
	return 0
}

func page56() int {
	fmt.Printf("That familiar strange feeling of deja'vu envelops you again.  It is hard to\n")
	fmt.Printf("say, but whatever is on the other side of the door does not seem to be intended\n")
	fmt.Printf("for you.\n")
	return choose(33, "You open the door and step through", 22, "You go looking for more information")
}

func page57() int {
	fmt.Printf("In the centre of the room is a table and a single chair.  There is an Orange\n")
	fmt.Printf("folder on the table top, but you can't make out the lettering on it.\n")
	return choose(11, "You sit down and read the folder", 12, "You leave the room")
}

func next_page(this_page int) int {
	fmt.Printf("\n")
	switch this_page {
	case 0:
		return 0
	case 1:
		return page1()
	case 2:
		return page2()
	case 3:
		return page3()
	case 4:
		return page4()
	case 5:
		return page5()
	case 6:
		return page6()
	case 7:
		return page7()
	case 8:
		return page8()
	case 9:
		return page9()
	case 10:
		return page10()
	case 11:
		return page11()
	case 12:
		return page12()
	case 13:
		return page13()
	case 14:
		return page14()
	case 15:
		return page15()
	case 16:
		return page16()
	case 17:
		return page17()
	case 18:
		return page18()
	case 19:
		return page19()
	case 20:
		return page20()
	case 21:
		return page21()
	case 22:
		return page22()
	case 23:
		return page23()
	case 24:
		return page24()
	case 25:
		return page25()
	case 26:
		return page26()
	case 27:
		return page27()
	case 28:
		return page28()
	case 29:
		return page29()
	case 30:
		return page30()
	case 31:
		return page31()
	case 32:
		return page32()
	case 33:
		return page33()
	case 34:
		return page34()
	case 35:
		return page35()
	case 36:
		return page36()
	case 37:
		return page37()
	case 38:
		return page38()
	case 39:
		return page39()
	case 40:
		return page40()
	case 41:
		return page41()
	case 42:
		return page42()
	case 43:
		return page43()
	case 44:
		return page44()
	case 45:
		return page45()
	case 46:
		return page46()
	case 47:
		return page47()
	case 48:
		return page48()
	case 49:
		return page49()
	case 50:
		return page50()
	case 51:
		return page51()
	case 52:
		return page52()
	case 53:
		return page53()
	case 54:
		return page54()
	case 55:
		return page55()
	case 56:
		return page56()
	case 57:
		return page57()
	default:
		return 0
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	instructions()
	more()
	character()
	more()
	for {
		page = next_page(page)
		if page == 0 {
			break
		}
		more()
	}
}
