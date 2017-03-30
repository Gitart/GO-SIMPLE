## Record voice(audio) from microphone to .WAV file

### Problem:
You need to activate your microphone, record your voice and save the data into a .wav file from command line. How to do that?

### Solution:
The code example below is combination of PortAudio (https://github.com/gordonklaus/portaudio/blob/master/examples/record.go)    
and Golang wave file reader/writer (https://github.com/zenwerk/go-wave/blob/master/example/writing.go).   
What it does is to activate your default microphone and start recording. 
Encode the data([]byte) to the wave.Write() function and save to a given .wav file.    

```golang
 package main

 /*
  #include <stdio.h>
  #include <unistd.h>
  #include <termios.h>
  char getch(){
      char ch = 0;
      struct termios old = {0};
      fflush(stdout);
      if( tcgetattr(0, &old) < 0 ) perror("tcsetattr()");
      old.c_lflag &= ~ICANON;
      old.c_lflag &= ~ECHO;
      old.c_cc[VMIN] = 1;
      old.c_cc[VTIME] = 0;
      if( tcsetattr(0, TCSANOW, &old) < 0 ) perror("tcsetattr ICANON");
      if( read(0, &ch,1) < 0 ) perror("read()");
      old.c_lflag |= ICANON;
      old.c_lflag |= ECHO;
      if(tcsetattr(0, TCSADRAIN, &old) < 0) perror("tcsetattr ~ICANON");
      return ch;
  }
 */
 import "C"

 // stackoverflow.com/questions/14094190/golang-function-similar-to-getchar

 import (
         "fmt"
         "github.com/gordonklaus/portaudio"
         wave "github.com/zenwerk/go-wave"
         "math/rand"
         "os"
         "strings"
         "time"
 )

 func errCheck(err error) {

         if err != nil {
                 panic(err)
         }
 }

 func main() {

         if len(os.Args) != 2 {
                 fmt.Printf("Usage : %s <audiofilename.wav>\n", os.Args[0])
                 os.Exit(0)
         }

         audioFileName := os.Args[1]

         fmt.Println("Recording. Press ESC to quit.")

         if !strings.HasSuffix(audioFileName, ".wav") {
                 audioFileName += ".wav"
         }
         waveFile, err := os.Create(audioFileName)
         errCheck(err)

         // www.people.csail.mit.edu/hubert/pyaudio/  - under the Record tab
         inputChannels := 1
         outputChannels := 0
         sampleRate := 44100
         framesPerBuffer := make([]byte, 64)

         // init PortAudio

         portaudio.Initialize()
         //defer portaudio.Terminate()

         stream, err := portaudio.OpenDefaultStream(inputChannels, outputChannels, float64(sampleRate), len(framesPerBuffer), framesPerBuffer)
         errCheck(err)
         //defer stream.Close()

         // setup Wave file writer

         param := wave.WriterParam{
                 Out:           waveFile,
                 Channel:       inputChannels,
                 SampleRate:    sampleRate,
                 BitsPerSample: 8, // if 16, change to WriteSample16()
         }

         waveWriter, err := wave.NewWriter(param)
         errCheck(err)

         //defer waveWriter.Close()

         go func() {
                 key := C.getch()
                 fmt.Println()
                 fmt.Println("Cleaning up ...")
                 if key == 27 {
                         // better to control
                         // how we close then relying on defer
                         waveWriter.Close()
                         stream.Close()
                         portaudio.Terminate()
                         fmt.Println("Play", audioFileName, "with a audio player to hear the result.")
                         os.Exit(0)

                 }

         }()

         // recording in progress ticker. From good old DOS days.
         ticker := []string{
                 "-",
                 "\\",
                 "/",
                 "|",
         }
         rand.Seed(time.Now().UnixNano())

         // start reading from microphone
         errCheck(stream.Start())
         for {
                 errCheck(stream.Read())

                 fmt.Printf("\rRecording is live now. Say something to your microphone! [%v]", ticker[rand.Intn(len(ticker)-1)])

                 // write to wave file
                 _, err := waveWriter.Write([]byte(framesPerBuffer)) // WriteSample16 for 16 bits
                 errCheck(err)
         }
         errCheck(stream.Stop())
 }
```

### References:

https://godoc.org/github.com/gordonklaus/portaudio
http://people.csail.mit.edu/hubert/pyaudio/

