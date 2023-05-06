package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randQuote() string {
	// set up random seed generation
	rand.Seed(time.Now().Unix())
	// set up array of quotes from Variks
	quotes := []string{
		"Fight. Win. Live.",
		"Fight. Kill. Survive.",
		"Impress me, yes? Fight and win!",
		"Impress me, yes?",
		"Prepare.",
		"Prepare for combat.",
		"They will want to kill you. Kill them back.",
		"They will try to kill you. Kill them back.",
		"Kill them dead, Guardian.",
		"Return/Retreat to airlock, Guardian...",
		"Go back to airlock, yes?",
		"Incoming...",
		"Reinforcements. Ready for combat.",
		"Reinforcements are coming.",
		"Incoming.",
		"Enemies advance.",
		"More enemies, Guardian.",
		"Be ready. More enemies on the way.",
		"More enemies have been unleashed.",
		"Dismantle mines, yes? Or... you die...",
		"Mines must be dismantled. Or death.",
		"Intercept important target. Or, you die.",
		"Success, Guardian, success.",
		"Success. You have done it.",
		"Fallen of the Wolf Banner march to war...",
		"You face Fallen pirate scum.",
		"Hive spawn claw their way towards your Light.",
		"Bonewalkers of the Hive await.",
		"Hive, creatures of the Darkness, arise.",
		"Foot soldiers of the Cabal rise to the challenge.",
		"Warbeasts of the Cabal await.",
		"Time lost Vex arise.",
		"Vex incursion. Meld of mind and machine.",
		"Servants of Oryx unleashed!",
		"Taken come, want your Light.",
		"Taken advance on your position.",
		"A servant of Oryx hungers for your Light.",
		"Taken approach, beware.",
		"Urrox the Flame Prince, scion of Oryx.",
		"Valus Trau'ug, Cabal Juggernaut.",
		"Cabal strategist Val Aru'un.",
		"The Wretched Knight. Dark blade of the Hive.",
		"The Overmind Minotaur.",
		"Pilot Servitor of a war-wracked Fallen Ketch.",
		"Keksis, the Betrayed!",
		"Noru'usk, Servant of Oryx.",
		"Sylok, the Defiled, thirsts for your Light.",
		"Victory. Well earned.",
		"Your victory tells the tale.",
		"Enemies die. You live. Well done.",
		"They are dead. You are not.",
		"You fought, you won. Your treasure awaits.",
		"Taste victory. It is sweet. You have won.",
	}
	// grab a random quote using our rand seed we created
	q := rand.Int() % len(quotes)
	return fmt.Sprintf("\n%v\n", quotes[q])
}

func main() {
	// grab a random variks quote
	vquote := randQuote()
	// send it to stdout
	fmt.Println(vquote)
}
