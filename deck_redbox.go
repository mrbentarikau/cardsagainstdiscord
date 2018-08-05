package cardsagainstdiscord

func init() {
	pack := &CardPack{
		Name:        "redbox",
		Description: "Red box expansion",
		Prompts: []*PromptCard{
			&PromptCard{Prompt: `%s would be woefully incomplete without %s.`},
			&PromptCard{Prompt: `%s: Hours of fun. Easy to use. Perfect for %s!`},
			&PromptCard{Prompt: `%s. Awesome in theory, kind of a mess in practice.`},
			&PromptCard{Prompt: `A remarkable new study has shown that chimps have evolved their own primitive version of %s.`},
			&PromptCard{Prompt: `A successful job interview begins with a firm handshake and ends with %s.`},
			&PromptCard{Prompt: `After months of practice with %s, I think I’m finally ready for %s.`},
			&PromptCard{Prompt: `And I would have gotten away with it, too, if it hadn't been for %s!`},
			&PromptCard{Prompt: `And what did you bring for show and tell?`},
			&PromptCard{Prompt: `As part of his contract, Prince won’t perform without %s in his dressing room.`},
			&PromptCard{Prompt: `As part of his daily regimen, Anderson Cooper sets aside 15 minutes for %s.`},
			&PromptCard{Prompt: `Awww, sick! I just say this skater do a 720 kickflip into %s!`},
			&PromptCard{Prompt: `Before %s, all we had was %s.`},
			&PromptCard{Prompt: `Before I run for president, I must destroy all evidence of my involvement with %s.`},
			&PromptCard{Prompt: `Call the law offices of Goldstein & Goldstein, because no one should have to tolerate %s in the workplace.`},
			&PromptCard{Prompt: `Charades was ruined for me forever when my mom had to act out %s.`},
			&PromptCard{Prompt: `Dear Sir or Madam, We regret to inform you that the Office of %s has denied your request for %s.`},
			&PromptCard{Prompt: `Doctor, you've gone too far! The human body wasn't meant to withstand that amount of %s!`},
			&PromptCard{Prompt: `During high school I never really fit in until I found %s club.`},
			&PromptCard{Prompt: `During his midlife crisis, my dad got really into %s.`},
			&PromptCard{Prompt: `Finally! A service that delivers %s right to your door.`},
			&PromptCard{Prompt: `Future Historians will agree that %s marked the beginning of America's decline.`},
			&PromptCard{Prompt: `Having problems with %s? Try %s!`},
			&PromptCard{Prompt: `Hey baby, come back to my place and I’ll show you %s.`},
			&PromptCard{Prompt: `I learned the hard way that you can't cheer up a grieving friend with %s.`},
			&PromptCard{Prompt: `I love being a mom. But it's tough when my kids come home filthy from %s. That's why there's Tide®.`},
			&PromptCard{Prompt: `I spent my whole life working toward %s, only to have it ruined by %s.`},
			&PromptCard{Prompt: `I went from %s to %s, all thanks to %s.`},
			&PromptCard{Prompt: `I’m not like the rest of you. I’m too rich and busy for %s.`},
			&PromptCard{Prompt: `If God didn't want us to enjoy %s, he wouldn't have given us %s.`},
			&PromptCard{Prompt: `In a pinch, %s can be a suitable substitute for %s.`},
			&PromptCard{Prompt: `In his new self-produced album, Kanye West raps over the sounds of %s.`},
			&PromptCard{Prompt: `In his newest and most difficult stunt, David Blaine must escape from %s.`},
			&PromptCard{Prompt: `In its new tourism campaign, Detroit proudly proclaims that it has finally eliminated %s.`},
			&PromptCard{Prompt: `In Rome, there are whisperings that the Vatican has a secret room devoted to %s.`},
			&PromptCard{Prompt: `In the seventh circle of Hell, sinners must endure %s for all eternity.`},
			&PromptCard{Prompt: `Legend has it Prince wouldn't perform without %s in his dressing room.`},
			&PromptCard{Prompt: `Listen, son. If you want to get involved with %s, I won’t stop you. Just steer clear of %s.`},
			&PromptCard{Prompt: `Little Miss Muffet Sat on a tuffet, Eating her curds and %s.`},
			&PromptCard{Prompt: `Lovin’ you is easy ’cause you’re %s.`},
			&PromptCard{Prompt: `Members of New York's social elite are paying thousands of dollars just to experience %s.`},
			&PromptCard{Prompt: `Michael Bay's new three-hour action epic pits %s against %s.`},
			&PromptCard{Prompt: `Money can’t buy me love, but it can buy me %s.`},
			&PromptCard{Prompt: `My country, 'tis of thee, sweet land of %s.`},
			&PromptCard{Prompt: `My gym teacher got fired for adding %s to the obstacle course.`},
			&PromptCard{Prompt: `My life is ruled by a vicious cycle of %s and %s.`},
			&PromptCard{Prompt: `My mom freaked out when she looked at my browser history and found %s.com/%s.`},
			&PromptCard{Prompt: `My new favorite porn star is Joey "%s" McGee.`},
			&PromptCard{Prompt: `My plan for world domination begins with %s.`},
			&PromptCard{Prompt: `Next time on Dr. Phil: How to talk to your child about %s.`},
			&PromptCard{Prompt: `Next week on the Discovery Channel, one man must survive in the depths of the Amazon with only %s and his wits.`},
			&PromptCard{Prompt: `Only two things in life are certain: death and %s.`},
			&PromptCard{Prompt: `Science will never explain %s.`},
			&PromptCard{Prompt: `The blind date was going horribly until we discovered our shared interest in %s.`},
			&PromptCard{Prompt: `The CIA now interrogates enemy agents by repeatedly subjecting them to %s.`},
			&PromptCard{Prompt: `The Five Stages of Grief: denial, anger, bargaining, %s, acceptance.`},
			&PromptCard{Prompt: `The healing process began when I joined a support group for victims of %s.`},
			&PromptCard{Prompt: `The secret to a lasting marriage is communication, communication, and %s.`},
			&PromptCard{Prompt: `This is your captain speaking. Fasten your seatbelts and prepare for %s.`},
			&PromptCard{Prompt: `This month's Cosmo: "Spice up your sex life by bringing %s into the bedroom."`},
			&PromptCard{Prompt: `To prepare for his upcoming role, Daniel Day-Lewis immersed himself in the world of %s.`},
			&PromptCard{Prompt: `Tonight on 20/20: What you don't know about %s could kill you.`},
			&PromptCard{Prompt: `Turns out that %s-Man was neither the hero we needed nor wanted.`},
			&PromptCard{Prompt: `Welcome to Señor Frog's! Would you like to try our signature cocktail, “%s on the Beach"?`},
			&PromptCard{Prompt: `What brought the orgy to a grinding halt?`},
			&PromptCard{Prompt: `What has been making life difficult at the nudist colony?`},
			&PromptCard{Prompt: `What left this stain on my couch?`},
			&PromptCard{Prompt: `What's harshing my mellow, man?`},
			&PromptCard{Prompt: `When all else fails, I can always masturbate to %s.`},
			&PromptCard{Prompt: `When I pooped, what came out of my butt?`},
			&PromptCard{Prompt: `When you get right down to it, %s is just %s.`},
			&PromptCard{Prompt: `With enough time and pressure, %s will turn into %s.`},
			&PromptCard{Prompt: `You haven't truly lived until you've experienced %s and %s at the same time.`},
			&PromptCard{Prompt: `Your persistence is admirable, my dear Prince. But you cannot win my heart with %s alone.`},
		},

		Responses: []ResponseCard{
			`24-hour media coverage.`,
			`A 55-gallon drum of lube.`,
			`A big black dick.`,
			`A bigger, blacker dick.`,
			`A black-owned and operated business.`,
			`A bloody pacifier.`,
			`A boo-boo.`,
			`A botched circumcision.`,
			`A Burmese tiger pit.`,
			`A cat video so cute that your eyes roll back and your spine slides out of your anus.`,
			`A cop who is also a dog.`,
			`A crappy little hand.`,
			`A dollop of sour cream.`,
			`A fat bald man from the internet.`,
			`A fortuitous turnip harvest.`,
			`A greased-up Matthew McConaughey.`,
			`A Japanese toaster you can fuck.`,
			`A Japanese tourist who wants something very badly but cannot communicate it.`,
			`A lamprey swimming up the toilet and latching onto your taint.`,
			`A low standard of living.`,
			`A magic hippie love cloud.`,
			`A man in yoga pants with a ponytail and feather earrings.`,
			`A nautical theme.`,
			`A nuanced critique.`,
			`A passionate Latino lover.`,
			`A phantasmagoria of anal delights.`,
			`A pile of squirming bodies.`,
			`A piñata full of scorpions.`,
			`A plunger to the face.`,
			`A PowerPoint presentation.`,
			`A sad fat dragon with no friends.`,
			`A sales team of clowns and pedophiles.`,
			`A slightly shittier parallel universe.`,
			`A smiling black man, a latina businesswoman, a cool Asian, and some whites.`,
			`A sofa that says “I have style, but I like to be comfortable.”`,
			`A spontaneous conga line.`,
			`A squadron of moles wearing aviator goggles.`,
			`A surprising amount of hair.`,
			`A sweaty, panting leather daddy.`,
			`A sweet spaceship.`,
			`A vagina that leads to another dimension.`,
			`A web of lies.`,
			`Actually getting shot, for real.`,
			`All my friends dying.`,
			`All of this blood.`,
			`An all-midget production of Shakespeare’s Richard III.`,
			`An army of skeletons.`,
			`An ass disaster.`,
			`An ether-soaked rag.`,
			`An Etsy steampunk strap-on.`,
			`An evil man in evil clothes.`,
			`An unhinged ferris wheel rolling toward the sea.`,
			`An unhinged Ferris wheel rolling toward the sea.`,
			`An unstoppable wave of fire ants.`,
			`André the Giant's enormous, leathery scrotum.`,
			`Another shot of morphine.`,
			`Basic human decency.`,
			`Being a busy adult with many important things to do.`,
			`Being a dinosaur.`,
			`Being a hideous beast that no one could love.`,
			`Being awesome at sex.`,
			`Being black.`,
			`Being white.`,
			`Big Bird's brown, crusty asshole.`,
			`Big ol' floppy titties.`,
			`Bill Clinton, naked on a bearskin rug with a saxophone.`,
			`Blood farts.`,
			`Blowing some dudes in an alley.`,
			`Boris the Soviet Love Hammer.`,
			`Bosnian chicken farmers.`,
			`Bullshit.`,
			`Buying the right pants to be cool.`,
			`Catastrophic urethral trauma.`,
			`Chugging a lava lamp.`,
			`Clams.`,
			`Clenched butt cheeks.`,
			`Cock.`,
			`Consent.`,
			`Converting to Islam.`,
			`Crabapples all over the fucking sidewalk.`,
			`Crushing Mr. Peanut's brittle body.`,
			`Crying into the pages of Sylvia Plath.`,
			`Cumming deep inside my best bro.`,
			`Dad's funny balls.`,
			`Daddy's belt.`,
			`Death by Steven Seagal.`,
			`Deflowering the princess.`,
			`Demonic possession.`,
			`Dining with cardboard cutouts of the cast of "Friends."`,
			`Dining with cardboard cutouts of the cast of Friends.`,
			`Disco fever.`,
			`Dorito breath.`,
			`Double penetration.`,
			`Drinking my bro's pee-pee right out of his peen.`,
			`Drinking ten 5-hour ENERGYs® to get fifty continuous hours of energy.`,
			`Dying alone and in pain.`,
			`Eating an albino.`,
			`Eating Tom Selleck’s mustache to gain his powers.`,
			`Emotional baggage.`,
			`Enormous Scandinavian women.`,
			`Existing.`,
			`Fabricating statistics.`,
			`Fetal alcohol syndrome.`,
			`Filling every orifice with butterscotch pudding.`,
			`Filling my son with spaghetti.`,
			`Finding Waldo.`,
			`Fisting.`,
			`Flying robots that kill people.`,
			`Fuck Mountain.`,
			`Gandalf.`,
			`Gargling jizz.`,
			`Gay aliens.`,
			`Genetically engineered super-soldiers.`,
			`George Clooney's musk.`,
			`Getting abducted by Peter Pan.`,
			`Getting hilariously gang-banged by the Blue Man Group.`,
			`Getting your dick stuck in a Chinese finger trap with another dick.`,
			`Gladiatorial combat.`,
			`Going around punching people.`,
			`Grandpa's ashes.`,
			`Graphic violence, adult language, and some sexual content.`,
			`Having $57 in the bank.`,
			`Having a penis.`,
			`Having sex on top of a pizza.`,
			`Having shotguns for legs.`,
			`Hillary Clinton.`,
			`Hipsters.`,
			`Historical revisionism.`,
			`Hot doooooooogs.`,
			`How awesome it is to be white.`,
			`How wet my pussy is.`,
			`Indescribable loneliness.`,
			`Insatiable bloodlust.`,
			`Intimacy problems.`,
			`Jafar.`,
			`Jeff Goldblum.`,
			`Jesus.`,
			`Jumping out at people.`,
			`Just the tip.`,
			`Letting everyone down.`,
			`Leveling up.`,
			`Literally eating shit.`,
			`Living in a trash can.`,
			`Living in a trashcan.`,
			`Loki, the trickster god.`,
			`Mad hacky-sack skills.`,
			`Making a friend.`,
			`Making shit up.`,
			`Making the penises kiss.`,
			`Maximal insertion.`,
			`Me.`,
			`Medieval Times® Dinner & Tournament.`,
			`Mild autism.`,
			`Mom.`,
			`Mooing.`,
			`Moral ambiguity.`,
			`Mufasa’s death scene.`,
			`Multiple orgasms.`,
			`My father, who died when I was seven.`,
			`My first kill.`,
			`My machete.`,
			`My manservant, Claude.`,
			`Neil Patrick Harris.`,
			`Not contributing to society in any meaningful way.`,
			`Not having sex.`,
			`Nothing.`,
			`Nubile slave boys.`,
			`Nunchuck moves.`,
			`Ominous background music.`,
			`Oncoming traffic.`,
			`One Ring to rule them all.`,
			`One thousand Slim Jims.`,
			`Overpowering your father.`,
			`Power.`,
			`Pretty Pretty Princess Dress-Up Board Game®.`,
			`Pumping out a baby every nine months.`,
			`Putting an entire peanut butter and jelly sandwich into the VCR.`,
			`Quiche.`,
			`Racial profiling.`,
			`Revenge fucking.`,
			`Reverse cowgirl.`,
			`Rich people.`,
			`Ripping into a man's chest and pulling out his still-beating heart.`,
			`Ripping open a man’s chest and pulling out his still-beating heart.`,
			`Rising from the grave.`,
			`Roland the Farter, flatulist to the king.`,
			`Running naked through a mall, pissing and shitting everywhere.`,
			`Ryan Gosling riding in on a white horse.`,
			`Samuel L. Jackson.`,
			`Sanding off a man's nose.`,
			`Santa Claus.`,
			`Savagely beating a mascot.`,
			`Screaming like a maniac.`,
			`Scrotal frostbite.`,
			`Scrotum tickling.`,
			`Self-flagellation.`,
			`Sexual humiliation.`,
			`Shaft.`,
			`Shitting out a screaming face.`,
			`Shutting the fuck up.`,
			`Slapping a racist old lady.`,
			`Sneezing, farting, and cumming at the same time.`,
			`Some douche with an acoustic guitar.`,
			`Some kind of bird-man.`,
			`Some really fucked-up shit.`,
			`Sorcery.`,
			`Special musical guest, Cher.`,
			`Spending lots of money.`,
			`Staring at a painting and going "hmmmmmm..."`,
			`Statistically validated stereotypes.`,
			`Stockholm Syndrome`,
			`Subduing a grizzly bear and making her your wife.`,
			`Sudden Poop Explosion Disease.`,
			`Suicidal thoughts.`,
			`Suicide bombers.`,
			`Survivor's guilt.`,
			`Swiftly achieving orgasm.`,
			`Syphilitic insanity.`,
			`Taking a man's eyes and balls out and putting his eyes where his balls go and then his balls in the eye holes.`,
			`That ass.`,
			`The baby that ruined my pussy.`,
			`The black Power Ranger.`,
			`The boners of the elderly.`,
			`The day the birds attacked.`,
			`The economy.`,
			`The entire Internet.`,
			`The flute.`,
			`The four arms of Vishnu.`,
			`The Gulags.`,
			`The Harlem Globetrotters.`,
			`The harsh light of day.`,
			`The hiccups.`,
			`The hose.`,
			`The human body.`,
			`The Land of Chocolate.`,
			`The mere concept of Applebee's®.`,
			`The milk that comes out of a person.`,
			`The mixing of the races.`,
			`The moist, demanding chasm of his mouth.`,
			`The new Radiohead album.`,
			`The ooze.`,
			`The primal, ball-slapping sex your parents are having right now.`,
			`The prunes I've been saving for you in my armpits.`,
			`The Quesadilla Explosion Salad™ from Chili’s.®`,
			`The shambling corpse of Larry King.`,
			`The systematic destruction of an entire people and their way of life.`,
			`The thin veneer of situational causality that underlies porn.`,
			`The total collapse of the global financial system.`,
			`The way white people is.`,
			`The worst pain imaginable. Times two!`,
			`Three months in the hole.`,
			`Tiny nipples.`,
			`Tongue.`,
			`Tripping balls.`,
			`Unlimited soup, salad, and breadsticks.`,
			`Velcro™.`,
			`Vietnam flashbacks.`,
			`Vomiting mid-blowjob.`,
			`Walking in on Dad peeing into Mom's mouth.`,
			`Warm, velvety muppet sex.`,
			`Weapons-grade plutonium.`,
			`Wearing an octopus for a hat.`,
			`Whining like a little bitch.`,
			`Whipping a disobedient slave.`,
			`White power.`,
			`Words, words, words.`,
			`Zeus's sexual appetites.`,
		},
	}

	AddPack("redbox", pack)
}
