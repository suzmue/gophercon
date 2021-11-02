package garden

import (
	"context"
	"fmt"
	"runtime"
	"strings"
)

var palmTree = "###########################'`################################\n" +
	"###########################  V##'############################\n" +
	"#########################V'  `V  ############################\n" +
	"########################V'      ,############################\n" +
	"#########`#############V      ,A###########################V\n" +
	"########' `###########V      ,###########################V',#\n" +
	"######V'   ###########l      ,####################V~~~~'',###\n" +
	"#####V'    ###########l      ##P' ###########V~~'   ,A#######\n" +
	"#####l      d#########l      V'  ,#######V~'       A#########\n" +
	"#####l      ##########l         ,####V''         ,###########\n" +
	"#####l        `V######l        ,###V'   .....;A##############\n" +
	"#####A,         `######A,     ,##V' ,A#######################\n" +
	"#######A,        `######A,    #V'  A########'''''##########''\n" +
	"##########,,,       `####A,           `#''           '''  ,,,\n" +
	"#############A,                               ,,,     ,######\n" +
	"######################oooo,                 ;####, ,#########\n" +
	"##################P'                   A,   ;#####V##########\n" +
	"#####P'    ''''       ,###             `#,     `V############\n" +
	"##P'                ,d###;              ##,       `V#########\n" +
	"##########A,,   #########A              )##,    ##A,..,ooA###\n" +
	"#############A, Y#########A,            )####, ,#############\n" +
	"###############A ############A,        ,###### ##############\n" +
	"###############################       ,#######V##############\n" +
	"###############################      ,#######################\n" +
	"##############################P    ,d########################\n" +
	"##############################'    d#########################\n" +
	"##############################     ##########################\n" +
	"##############################     ##########################\n" +
	"#############################P     ##########################\n" +
	"#############################'     ##########################\n" +
	"############################P      ##########################\n" +
	"###########################P'     ;##########################\n" +
	"###########################'     ,###########################\n" +
	"##########################       ############################\n" +
	"#########################       ,############################\n" +
	"########################        d###########P'    `Y#########\n" +
	"#######################        ,############        #########\n" +
	"######################        ,#############        #########\n" +
	"#####################        ,##############b.    ,d#########\n" +
	"####################        ,################################\n" +
	"###################         #################################\n" +
	"##################          #######################P'  `V##P'\n" +
	"#######P'     `V#           ###################P'\n" +
	"#####P'                    ,#################P'\n" +
	"###P'                      d##############P''\n" +
	"##P'                       V##############'\n" +
	"#P'                         `V###########'\n" +
	"#'                             `V##P'\n" +
	"\n" +
	"                                                        GNN94\n"

var route = getPineapple()

func getPineapple() []string {
	return strings.Split(palmTree, "\n")
}

func unlockDoor(passcode string) bool {
	return passcode == "palm tree"
}

func flap(i int) string {
	if i%(len(route)*2) == 0 {
		runtime.Gosched()
	}
	x := route[i%len(route)]
	return x
}

var quit = make(chan bool)

func openDoor() {
	defer func() {
		for i := 0; i < NUM_BUTTERFLIES; i++ {
			quit <- true
		}
	}()
	fmt.Println("A map of the island! There's an X on the beach under the pier. (Go to the beach and set a breakpoint on beach.Pier to find the spot to dig.")
}

func butterflies(ctx context.Context) {
	for i := 0; i < NUM_BUTTERFLIES; i++ {
		go func(butterflyId int) {
			for j := 0; ; j++ {
				select {
				case <-ctx.Done():
					return // returning not to leak the goroutine
				default:
					loc := flap(j)
					_ = watchButterflies(butterflyId, loc)
				}
			}
		}(i)
	}
}