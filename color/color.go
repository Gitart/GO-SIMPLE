package main
import  "fmt"
import "flag"
// import  "app/build"


// Structure
type Admin struct {
     Name   string
     Title  string
     Num    int64
}
var Version = "development"


func (r *Admin) Adm() string {
      return r.Name + " " + r.Title
}

func (r *Admin) Test() string {
      return r.Name + " " + r.Title
}



// Color site information
// http://www.lihaoyi.com/post/BuildyourownCommandLinewithANSIescapecodes.html
// https://stackoverflow.com/questions/5762491/how-to-print-color-in-console-using-system-out-println
// https://stackoverflow.com/questions/5762491/how-to-print-color-in-console-using-system-out-println
type Color string

const (
    ColorBlack  Color         = "\u001b[30m"
    ColorRed                  = "\u001b[31m"
    ColorGreen                = "\u001b[32m"
    ColorYellow               = "\u001b[33m"
    ColorBlue                 = "\u001b[34m"
    ColorPurpur               = "\u001b[35m"
    ColorPur                  = "\u001b[36m"
    ColorWhite                = "\u001b[37m"
    ColorP                    = "\u001b[29m"
    ColorReset                = "\u001b[0m"

    // High Intensity backgrounds
    BLACK_BACKGROUND_BRIGHT   = "\033[0;100m";  // BLACK
    RED_BACKGROUND_BRIGHT     = "\033[0;101m";  // RED
    GREEN_BACKGROUND_BRIGHT   = "\033[0;102m";  // GREEN
    YELLOW_BACKGROUND_BRIGHT  = "\033[0;103m";  // YELLOW
    BLUE_BACKGROUND_BRIGHT    = "\033[0;104m";  // BLUE
    PURPLE_BACKGROUND_BRIGHT  = "\033[0;105m";  // PURPLE
    CYAN_BACKGROUND_BRIGHT    = "\033[0;106m";  // CYAN
    WHITE_BACKGROUND_BRIGHT   = "\033[0;107m";  // WHITE

    // Bold High Intensity
     BLACK_BOLD_BRIGHT        = "\033[1;90m";   // BLACK
     RED_BOLD_BRIGHT          = "\033[1;91m";   // RED
     GREEN_BOLD_BRIGHT        = "\033[1;92m";   // GREEN
     YELLOW_BOLD_BRIGHT       = "\033[1;93m";   // YELLOW
     BLUE_BOLD_BRIGHT         = "\033[1;94m";   // BLUE
     PURPLE_BOLD_BRIGHT       = "\033[1;95m";   // PURPLE
     CYAN_BOLD_BRIGHT         = "\033[1;96m";   // CYAN
     WHITE_BOLD_BRIGHT        = "\033[1;97m";   // WHITE

    // Regular Colors
     BLACK                    = "\033[0;30m";   // BLACK
     RED                      = "\033[0;31m";   // RED
     GREEN                    = "\033[0;32m";   // GREEN
     YELLOW                   = "\033[0;33m";   // YELLOW
     BLUE                     = "\033[0;34m";   // BLUE
     PURPLE                   = "\033[0;35m";   // PURPLE
     CYAN                     = "\033[0;36m";   // CYAN
     WHITE                    = "\033[0;37m";   // WHITE

    // Bold
     BLACK_BOLD               = "\033[1;30m";   // BLACK
     RED_BOLD                 = "\033[1;31m";   // RED
     GREEN_BOLD               = "\033[1;32m";   // GREEN
     YELLOW_BOLD              = "\033[1;33m";   // YELLOW
     BLUE_BOLD                = "\033[1;34m";   // BLUE
     PURPLE_BOLD              = "\033[1;35m";   // PURPLE
     CYAN_BOLD                = "\033[1;36m";   // CYAN
     WHITE_BOLD               = "\033[1;37m";   // WHITE

    // Underline
     BLACK_UNDERLINED         = "\033[4;30m";   // BLACK
     RED_UNDERLINED           = "\033[4;31m";   // RED
     GREEN_UNDERLINED         = "\033[4;32m";   // GREEN
     YELLOW_UNDERLINED        = "\033[4;33m";   // YELLOW
     BLUE_UNDERLINED          = "\033[4;34m";   // BLUE
     PURPLE_UNDERLINED        = "\033[4;35m";   // PURPLE
     CYAN_UNDERLINED          = "\033[4;36m";   // CYAN
     WHITE_UNDERLINED         = "\033[4;37m";   // WHITE

    // Background
     BLACK_BACKGROUND         = "\033[40m";     // BLACK
     RED_BACKGROUND           = "\033[41m";     // RED
     GREEN_BACKGROUND         = "\033[42m";     // GREEN
     YELLOW_BACKGROUND        = "\033[43m";     // YELLOW
     BLUE_BACKGROUND          = "\033[44m";     // BLUE
     PURPLE_BACKGROUND        = "\033[45m";     // PURPLE
     CYAN_BACKGROUND          = "\033[46m";     // CYAN
     WHITE_BACKGROUND         = "\033[47m";     // WHITE

    // High Intensity
     BLACK_BRIGHT             = "\033[0;90m";   // BLACK
     RED_BRIGHT               = "\033[0;91m";   // RED
     GREEN_BRIGHT             = "\033[0;92m";   // GREEN
     YELLOW_BRIGHT            = "\033[0;93m";   // YELLOW
     BLUE_BRIGHT              = "\033[0;94m";   // BLUE
     PURPLE_BRIGHT            = "\033[0;95m";   // PURPLE
     CYAN_BRIGHT              = "\033[0;96m";   // CYAN
     WHITE_BRIGHT             = "\033[0;97m";   // WHITE
)

func colorize(color Color, message string) {
     fmt.Println(string(color), message, string(ColorReset))
}


func clRed(Message string){
     s:=string(ColorReset)
     c:=string(RED_BOLD_BRIGHT)
     fmt.Println(c,Message, s)	
}

func clGreen(Message string){
     s:=string(ColorReset)
     c:=string(GREEN_BOLD_BRIGHT)
     fmt.Println(c,Message, s)	
}


// Двойной цвет
func clGreenRed(Message,Value string){
     s:=string(ColorReset)
     c:=string(GREEN_BOLD_BRIGHT)
     y:=string(YELLOW_BOLD_BRIGHT)
     fmt.Println(c,Message,y, Value,  s)	
}


func main1() {
    useColor := flag.Bool("color", false, "display colorized output")
    flag.Parse()

    if *useColor {
        colorize(ColorBlue, "Hello, DigitalOcean!")
        return
    }
}


// Color to command line
func Colors(){
	colorize(RED_BACKGROUND_BRIGHT,      "Hello, DigitalOcean!")
    clRed("Ysee")
    clGreen("Ysee")
    clGreenRed("Пример","122")
    // colorize(ColorBlue,   "Hello, DigitalOcean!")
    // colorize(ColorRed,    "Hello, DigitalOcean!")
    // colorize(ColorGreen,  "Hello, DigitalOcean!")
    // colorize(ColorYellow, "Hello, DigitalOcean!")
    fmt.Println("Hello, DigitalOcean!")
}


// Main
func main(){
	Colors()
	fmt.Println("Version:\t", Version)
    // main1()
    // f3()
}
