package main

import (
	"math/rand"
	"net/url"
	"os"
	"time"

	"github.com/ChimeraCoder/anaconda"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/sirupsen/logrus"
)

var (
	consumerKey       = getenv("TWITTER_CONSUMER_KEY")
	consumerSecret    = getenv("TWITTER_CONSUMER_SECRET")
	accessToken       = getenv("TWITTER_ACCESS_TOKEN")
	accessTokenSecret = getenv("TWITTER_ACCESS_TOKEN_SECRET")

	log = &logger{logrus.New()}
)

func getenv(name string) string {
	v := os.Getenv(name)
	if v == "" {
		panic("did you forget your keys? " + name)
	}
	return v
}

func lyric() {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

	api.SetLogger(log)

	lyrics := []string{
		"From the black bag skip in the parking lot, It's a short bad trip to the candy shop",
		"Where the shrimps sell smack to the jelly snakes, And the kids buy crack in their morning break",
		"And the grass grows bluer on the other side, Where the old girls queue for their Mothers Pride",
		"For a slice of life it's a bargain sale, The price is right but the bread is stale",
		"From the high rise priest of the office blocks, To a five year lease on a cardboard box",
		"From the Old Queens Head to the Burger King, In my '57 Chevvy made from baked bean tins",
		"And when I drive that heap down the road, You can hear the cheap car stereo",
		"Volume knob turned down low, And rubbish on the radio",
		"R U B B I S H",
		"I'm underage and uninsured, On the High Road to Domestos",
		"Chloraflouracarbon Lord, Asbestos lead asbestos",
		"I'm going now to meet my maker, Ignition, clutch, accelerator",
		"Mirror, signal and manoeuvre, Glory, glory, glory hallelujah!",
		"Bloodsport for all, said Corporal Flash, And shoved me in a room full of C.S. gas",
		"Stand up and beg, said Sergeant Kirby, Lay down, play dead for Di and Fergie",
		"Roll up, roll up goes the reveille, Abuse the bugle boy of Company B",
		"And the coldest stream guards of them all, Sang God Save The Queen and Bloodsport For All",
		"Both my arms and legs, Are torn to shreds, My eyes are tired and grey",
		"I've lost a stone, I'm just skin and bone, And when I come home today, Look away",
		"I'm a G.I. and I'm blue",
		"Look away John F Kennedy",
		"Look away Franklin D Roosevelt",
		"Look away George Washington, Thomas Jefferson and Brother Jonathon",
		"Look away Bob Hope",
		"Look away Uncle Sam",
		"Look away Ronald Reagan",
		"Look away Dixieland",
		"Oh God I wish I was in Dixie, In Dixie, with you",
		"When bonny Clive was twenty three, He took a dive from the balcony",
		"Well I remember Micky Doyle, He shuffled off this mortal coil",
		"Give me the beat boy and free my soul, Fill my pockets up with gold",
		"I'll leave a message on the fridge, And drive my car off London Bridge",
		"So I'm canceling my driving test, And walking back to happiness",
		"Whoop bye oh yeah yeah!",
		"Everytime a churchbell rings, Another angel gets its wings",
		"If your conscience fails you we can be your guide",
		"Johnny Guitar, tell 'em where it goes",
		"We've got smackheads, crackheads, pensioners, pimps, Anonymous alcoholics looking for a drink",
		"Twenty four minutes from Tulse Hill let's go",
		"We've got yardies, steamers, parasitic cops, Bostik boys playing chicken in the box,",
		"Mad alsations, pit-bull terrorists, Hammerheaded loan sharks trying out for Jaws 6",
		"You don't need a weatherman to know which way the wind blows",
		"Twenty four minutes from Tulse Hill, The driver's dressed in black, He's dead on the dead man's handle, And we ain't coming back",
		"We're going down the tracks and off the page, Past the dole, The Silver Blades",
		"Awopbopaloobopalopbamboom!",
		"Calling all cars, calling all cars, Check all the pubs and raid all the bars",
		"Fun, fun, fun, Here we come!",
		"Now, Sheriff Fatman started out in business as a granny farmer",
		"He was infamous for fifteen minutes and he appeared on Panorama",
		"Moving up on second base, Behind Nicholas Van Wotsisface",
		"At six foot six and a hundred tons, The undisputed king of the slums",
		"With more aliases than Klaus Barbie, The master butcher of Leigh-on-Sea",
		"Just about to take the stage, The one-and-only, hold the front page",
		"There's bats in the belfry, The windows are jammed, The toilets ain't healthy, And he don't give a damn",
		"He's buying up houses for the has-beens and healthy, From Lands End to Southend and Chelsea",
		"It was midnight on the murder mile, Wilson Pickett's finest hour",
		"From the gas board to the fire brigade, There's a dozen GPO's, An all night chicken takeaway, Which was finger lickin' closed",
		"It was midnight on the murder mile, O.K. let's riot",
		"In the avenues and alleyways, I took a short-cut to the throat",
		"I was stitched up by the boys brigade, And I was beaten to a pulp",
		"If the concrete and the clay beneath your feet, Don't get you son, The avenues and alleyways are gonna do it, Just for fun",
		"The telephones on sticks will tell you 999 calls only",
		"I need communion, confirmation and absolution for my crimes",
		"I need a character witness Jesus I think I'm about to die",
		"A public execution that the whole neighbourhood could watch",
		"Or just a phone box, a phone box, my kingdom for a phone box",
		"To the North of Katmandu, There's tiny children sniffing glue",
		"Like losers in a Michael Winner script, It's got a BMX certificate",
		"Outside there's a blazing sun, It's a perfect day to drop the bomb",
		"I James Robert Injustice, being of unsound body and mind, hereby bequeath all my worldly goods to anyone who wants 'em.",
		"Give my body to medical science, If medical science will have me",
		"They can take my lungs and kidneys, But my heart belongs to Daphne",
		"Are you prepared to meet your maker, And ask for your money back?",
		"This is my second to last will and testament, Only a rough draft, A handwritten estimate",
		"Don't bury me at sea, The pollution might kill me!",
		"The tequila sun is rising, And the Harvey's Bristol moon is sinking",
		"Put the Binatone on snooze, Open up some Special Brews and start drinking",
		"Yes sir, the Thunderbirds have gone, And the wagon's rolling on and I'm on it",
		"Disneyland or dipsomania, Pick a flavour",
		"Anytime, anyplace, anywhere, There's a wonderful world you can share",
		"Try agrophobic, schizophrenic, Paranoid attacks of panic, Epileptic fit of laughter, Twenty five million mornings after",
		"Moonshine, firewater, Captain Morgan, Johnnie Walker",
		"Southern Comfort, mother's ruin, Happy hours of homeless brewing",
		"Galloway's sore throat expectorant, After-shave and disinfectant",
		"Parazone and Fairy Liquid, If it's in a glass, you'll drink it",
		"Ground floor, shoppers' paradise, Haberdashery, needles, spoons, and knives",
		"Spend your money, girls, on sprays and lipsticks, Tested on bunnies, girls, strays, and misfits",
		"We've got a free pair of flares with every hip replacement, Just take the stairs to the bargain basement",
		"The big shop is open, It's a wonderful world",
		"Going down for all the things you missed, All the love, peace, and happiness that don't exist",
		"We've got encylopaedias, we've got pic 'n' fix, A government freezer full of benefits",
		"A children's assortment, we're bigger than Hamley's, We've got Cabbage Patch orphans from Sylvanian Families",
		"Carpets, linoleum, holy petroleum, Chemi-kaze killers, little Hitlers and Napoleons",
		"Ladies and gentlemen, boys and girls, The big shop is open and the world is Wonderful",
		"The great cucumber robberies of 1989",
		"Love ain't like the movies, It blisters and bruises, And knocks you around with its fists",
		"It leaves you a wreckage, all postage and package, Sealed with Glasgow kiss",
		"You win some, And you lose some, But I've lost the will to lose, My part time job, My faith in God, Falling on a bruise",
		"Falling on a bruise",
		"You need your Nutrasweet Daddy, Or some Peppermint Patty, Or just a hackneyed old cabby, Who can take you and your baby away",
		"It's not that I'm agorophobic... it's just that it's not safe to go outside anymore.",
		"I've looted and I've begged, on the tubes of the Bec, and Broadway",
		"I've been run over by cars, And to prove it here's the scars, On my wrist",
		"I've been cut, I've been stitched, I've been buggered, bewitched, and abadoned",
		"The Final Comedown, It's a victory worth sharing",
		"Hello, good evening and welcome, To nothing much",
		"The comfort and the joy, Of feeling lost, With the only living boy In New Cross",
		"I've teamed up with the hippies now, I've got my fringe unfurled",
		"I want to give peace, love and kisses out, To this whole stinking world",
		"The good, the bad, the average and unique",
		"The grebos, the crusties, the goths",
		"Because you're popular, And you're beautiful, The whole world knows your name",
		"But suppose you gave a funeral, And nobody came?",
		"Here's the church, Here's the steeple, Open the doors, Where's all the people?",
		"Between the Open University and closedown, You were dead",
		"John Player Special Number 666",
		"A black eye for a black eye, A chipped tooth for a chipped tooth",
		"A fraction of a half life, Some housework and some home truths, And nothing but the home truth",
		"And it's goodbye Ruby Tuesday, Come home you silly cow, We've baked a cake and your friends are waiting",
		"And David Icke says he'd like to show us how, To love you back to life again now",
		"He flies through the air, With the greatest of ease, That daring young man, In the blue dungarees",
		"And his poor, pathetic parents, So stricken with grief, That they've spelt his name wrong, On his funeral wreath!",
		"Yesterday they took away our bus stop, Today they'll try and take our happy home",
		"If we club together, with all the diamonds we've saved, we could look to our hearts and say, we've got it in spades.",
		"The ambulance sirens rang, as they wheeled her to the stand",
		"Did the three bears shit in the woods?",
		"The wheels of justice turned, 'til the tyres were bald and burned",
		"I'm underage and uninsured, On the High Road to Domestos",
	}

	rand.Seed(time.Now().Unix())
	n := rand.Int() % len(lyrics)

	_, err := api.PostTweet(lyrics[n], url.Values{})
	if err != nil {
		log.Critical(err)
	}
}

func main() {
	lambda.Start(lyric)
}

type logger struct {
	*logrus.Logger
}

func (log *logger) Critical(args ...interface{})                 { log.Error(args...) }
func (log *logger) Criticalf(format string, args ...interface{}) { log.Errorf(format, args...) }
func (log *logger) Notice(args ...interface{})                   { log.Info(args...) }
func (log *logger) Noticef(format string, args ...interface{})   { log.Infof(format, args...) }
