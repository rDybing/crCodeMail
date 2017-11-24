/*********************************
 * Written in Go 1.9.1 linux/amd64
 * though any 1.x.x version should
 * work fine under any OS.
 *
 * CC-BY Roy Dybing 2017
 * github: rDybing
 * slack : rdybing
 **********************************/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var quit bool
	var message []string
	help()
	for quit == false {
		in, echo := getInput()
		switch in {
		case "":
			help()
		case "quit":
			quit = true
		default:
			key := keyGen(in)
			/*
				message = dataStringClear()
				for i := range message {
					message[i] = encrypt(key, message[i])
					fmt.Println("data = append(data, " + message[i] + ")")
				}
			*/
			message = dataStringEncoded()
			if decrypt(key, message[0]) != "Greetings" {
				fmt.Printf("The force is weak in %s\n", echo)
			} else {
				fmt.Println()
				for i := range message {
					message[i] = Decrypt(key, message[i])
					fmt.Println(message[i])
				}
			}
		}
	}
	os.Exit(0)
}

func help() {
	fmt.Println("-------------------")
	fmt.Println("--- Coder Radio ---")
	fmt.Println("-------------------")
	fmt.Println("- Type quit and then return to exit")
	fmt.Println("- Empty string will repeat this message")
	fmt.Println("- Anything else will be treated as the decoding key")
	fmt.Println("  (hint: It's Mike's all time favourite")
	fmt.Println("  Star Wars character - full name, no slacking!)")
}

func getInput() (string, string) {
	fmt.Print("-> ")
	input := bufio.NewReader(os.Stdin)
	line, _ := input.ReadString('\n')
	unformatted := line
	line = strings.ToLower(line)
	line = strings.Replace(line, " ", "", -1)
	line = strings.Replace(line, "\n", "", -1)
	return line, unformatted
}

func keyGen(in string) string {
	var key int
	key = 2
	for i, r := range in {
		key *= int(r) ^ i
	}
	keyString := strconv.Itoa(key)
	return keyString
}

func encrypt(key string, in string) string {

	var temp int
	var out string
	pLen := len(in) - 1
	keyLength := len(key) - 1
	keyCount := 0

	for i := 0; i <= pLen; i++ {
		temp = int(in[i])
		if keyCount > keyLength {
			keyCount = 0
		}
		temp += int(key[keyCount])
		keyCount++
		out += strconv.Itoa(temp) + ","
	}
	return out
}

func decrypt(key string, in string) string {
	var out string
	var inTemp int
	var keyValue int
	keyLength := len(key) - 1
	keyCount := 0
	ts := strings.Split(in, ",")

	for i := 0; i < strings.Count(in, ","); i++ {
		if keyCount > keyLength {
			keyCount = 0
		}
		inTemp, _ = strconv.Atoi(ts[i])
		keyValue = int(key[keyCount])
		keyCount++
		out += string(inTemp - keyValue)
	}
	return out
}

func dataStringClear() []string {
	var data []string
	data = append(data, "Greetings")
	data = append(data, "From                                             zr")
	data = append(data, "Norway                                         Yt$$$.")
	data = append(data, "                                           .,e$$$$$F'")
	data = append(data, "                         4e r               $$$$$$$.")
	data = append(data, "                         d$$br            _z$$$$$$$F`")
	data = append(data, "                          ?$$b._          ^?$$$$$$$")
	data = append(data, "                           4$$$^     -eec.  ^^JP^ ..eee$%..")
	data = append(data, "                           -**N #c   -^***.eE  ^z$P$$$$$$$$$r-")
	data = append(data, "                  .ze$$$$$eu?$eu '$$$$b $=^*$$ .$$$$$$$$$$^")
	data = append(data, "               --.^?$$$$$$$$$c^$$c .^^^ e$K  =^^*?$$$P^^^^")
	data = append(data, "   ueee. `:`  $E !!h ?$$$$$$$$b R$N'~!! *$$F J^^^C.  `")
	data = append(data, "  J  `^$$eu`!h !!!`4!!<?$$$$$$$P ?^.eee-z.ee^ ~$$e.br")
	data = append(data, "  'j$$Ne`?$$c`4!~`-e-:!:`$$$$$$$ $$**^z $^R$P  3 ^$$$bJ")
	data = append(data, "   4$$$F^.`?$$c`!! V].!!!`?$$$$F.$$$# $u$% ee*^^ :4`^$^?$q")
	data = append(data, "    ^^`,!!!:`$$N.4!!~~.~~4 ?$$F'$$F.@.* -L.e@$$$$ec.      ^")
	data = append(data, "    ^Rr`!!!!h ?$$c`h: `# !! $F,r4$L***  e$$$$$$$$$$$$hc")
	data = append(data, "      #e'4!!!!L`$$b'!.:!h`~~ .$F'^    d$$$$$$$$$$$$$$$$$h,")
	data = append(data, "       ^$.`!!!!h $$b`!. -    $P /'   .$$$$$$$$$$$$$$$$$$$$$c")
	data = append(data, "         ^$c`!!!h`$$.4~      $$$r'   <$$$$$$$$$$$$$$$$$$$P^^^")
	data = append(data, "           ^te.`~ $$b        `Fue-   `$$$$$$$$$$$$$$P^.:  !! ^<")
	data = append(data, "              ^^=4$P^     .,,,. -^.   ?$$$$$$$$$$^?:. !! :!!~ ,,ec..")
	data = append(data, "                    ..z$$$$$$$$$h,    `$$$$$$P^..`!f :!f ~)Lze$$$P^^^^?i")
	data = append(data, "                  ud$$$$$$$$$$$$$$h    `?$$F <!!'<!>:~)ue$$P*^..:!!!!! J")
	data = append(data, "                .K$$$$$$$$$$$$$$$$$,     P.>e'!f !~ ed$$P^.:!!!!!!!!`.d^")
	data = append(data, "               z$$$$$$$$$$$$$$$$$$$$      4!!~Ve$$$P`:!!!!!!!!!!'.eP'")
	data = append(data, "              -*^. . ^??$$$$$$$$$$$$       ~ `z$$$F^.`!!!!!!!!!!',dP^")
	data = append(data, "            .^ ):!!h i`!- (^?$$$$$$f        ,$$P^:! ). `'!!!!`,d$F'")
	data = append(data, "       .ueeeu.J`-^.!h <-  ~`.. ??$$'       ,$$ :!!`e$$$$e `,e$F'")
	data = append(data, "    e$$$$$$$$$$$$$eeiC ^)?:-<:%':^?        ?$f !!! ?$$$$^,F^")
	data = append(data, "   P^....```^^?$$$$$$$$$euL^.!..` .         ^Tu._.,``^^")
	data = append(data, "   $ !!!!!!!!!!::.^^??$$$$$$eJ~^=.            ````")
	data = append(data, "   ?$.`!!!!!!!!!!!!!!:.^??$$$$$c'.")
	data = append(data, "    ^?b.`!!!!!!!!!!!!!!!!>.^?$$$$c")
	data = append(data, "      ^?$c`'!!!!!!!!!!!',eeb, ^$$$k")
	data = append(data, "         ^?$e.`'!!!!!!! $$$$$ ;.?$$")
	data = append(data, "            ^?$ee,``''!.^?$P`i!! 3P")
	data = append(data, "                ^^??$bec,,.,ceeeP^")
	data = append(data, "                       `^^^^^^")
	data = append(data, " ")
	data = append(data, "Merry Christmas and may January 1st be survivable if painful :)")
	data = append(data, " ")
	data = append(data, "Now then, for the question at hand - but first a bit of background.")
	data = append(data, "About a year ago I got onboard as head tech-potato at a local small")
	data = append(data, "company who provides a platform for teaching those with hearing and/or")
	data = append(data, "speech impairment to learn sign-language. And more important, that those")
	data = append(data, "close to the impaired may keep up with the learning process and build")
	data = append(data, "their vocabulary of signs in sync with the impaired.")
	data = append(data, " ")
	data = append(data, "At the time I got onboard, we already had an app out on iOS and Android")
	data = append(data, "that have been a limited success. It do its' job by all means, but is a")
	data = append(data, "bit clunky. And also native, meaning two code-bases to keep track of,")
	data = append(data, "in addition to the Java backend which also have some issues. All this is")
	data = append(data, "still live, with some minor fixes here and there. Mostly having made")
	data = append(data, "queries to the DB more optimized for better response times")
	data = append(data, " ")
	data = append(data, "Anyhow, this all was targeted directly to consumers - but we soon found")
	data = append(data, "that to make it a viable business, we had to target municipalities,")
	data = append(data, "institutions and other support network. And get it ready for the new")
	data = append(data, "Privacy and Security Regulations that kick into effect here in Europe")
	data = append(data, "next year.")
	data = append(data, " ")
	data = append(data, "A lot of planning, a lot of new features, and also a new target platform;")
	data = append(data, "the browser by means of a web-app in addition to iOS and Android.")
	data = append(data, " ")
	data = append(data, "In short, a complete rewrite. Greenfield as far as the horizon span! This")
	data = append(data, "time none of that native nonsense. We just can't keep up with three frontend")
	data = append(data, "codebases *and* backend. Our resources are limited.")
	data = append(data, " ")
	data = append(data, "Our alternatives soon boiled down to two: C# + Xamarin, or a more")
	data = append(data, "traditional web-stack using Nativescript + Typescript + Angular4. Oh, and")
	data = append(data, "C# + dotnet core + redis + postgres in the rear.")
	return data
}

func dataStringEncoded() []string {
	var data []string
	data = append(data, "116,166,153,156,168,156,166,152,165,")
	data = append(data, "115,166,163,164,84,83,88,81,82,81,86,87,86,84,87,82,83,82,80,80,77,84,84,87,84,83,88,81,82,81,86,87,86,84,87,82,83,82,80,80,77,84,84,87,84,83,88,81,82,171,168,")
	data = append(data, "123,163,166,174,149,172,88,81,82,81,86,87,86,84,87,82,83,82,80,80,77,84,84,87,84,83,88,81,82,81,86,87,86,84,87,82,83,82,80,80,77,84,84,87,84,83,88,138,166,85,90,91,100,")
	data = append(data, "77,84,84,87,84,83,88,81,82,81,86,87,86,84,87,82,83,82,80,80,77,84,84,87,84,83,88,81,82,81,86,87,86,84,87,82,83,82,80,80,77,84,84,101,96,152,92,85,86,85,90,125,93,")
	data = append(data, "77,84,84,87,84,83,88,81,82,81,86,87,86,84,87,82,83,82,80,80,77,84,84,87,84,103,157,81,164,81,86,87,86,84,87,82,83,82,80,80,77,84,84,87,88,87,92,85,86,85,90,101,")
	data = append(data, "77,84,84,87,84,83,88,81,82,81,86,87,86,84,87,82,83,82,80,80,77,84,84,87,84,151,92,85,148,163,86,87,86,84,87,82,83,82,80,80,77,84,147,177,88,87,92,85,86,85,90,125,150,")
	data = append(data, "77,84,84,87,84,83,88,81,82,81,86,87,86,84,87,82,83,82,80,80,77,84,84,87,84,83,119,85,86,147,100,150,86,84,87,82,83,82,80,80,77,84,146,118,88,87,92,85,86,85,90,")
	data = append(data, "77,84,84,87,84,83,88,81,82,81,86,87,86,84,87,82,83,82,80,80,77,84,84,87,84,83,88,101,86,85,90,149,86,84,87,82,83,95,149,149,144,98,84,87,146,145,130,129,144,81,100,101,155,153,156,86,88,96,94,")
	data = append(data, "77,84,84,87,84,83,88,81,82,81,86,87,86,84,87,82,83,82,80,80,77,84,84,87,84,83,88,94,92,91,132,87,89,151,87,82,83,95,142,90,87,94,98,156,121,83,88,143,172,85,134,91,90,88,91,86,87,86,84,84,159,97,")
	data = append(data, "77,84,84,87,84,83,88,81,82,81,86,87,86,84,87,82,83,82,94,170,146,88,88,91,88,87,157,166,113,85,155,172,86,91,91,86,87,86,146,80,81,113,146,97,88,87,88,95,86,85,90,91,90,88,91,86,87,86,142,")
	data = append(data, "77,84,84,87,84,83,88,81,82,81,86,87,86,84,87,95,96,96,142,111,81,88,88,91,88,87,92,85,86,148,148,91,90,151,87,96,145,144,142,80,146,88,127,87,84,112,150,143,92,112,90,91,90,132,149,144,145,144,")
	data = append(data, "77,84,84,172,153,152,157,95,82,145,112,151,86,84,91,119,83,83,81,152,77,115,88,91,88,87,92,85,86,85,152,87,136,88,133,89,177,83,81,80,87,88,88,125,84,125,150,143,144,116,100,87,86,148,")
	data = append(data, "77,84,126,87,84,147,150,85,86,150,171,151,87,156,87,83,84,83,144,100,78,85,112,118,88,87,92,85,86,85,90,135,86,115,149,96,152,151,149,93,167,98,153,156,146,83,182,85,86,150,100,153,168,")
	data = append(data, "77,84,91,161,88,87,134,150,146,112,90,91,153,148,107,83,177,146,93,149,90,110,85,113,148,87,92,85,86,85,90,91,86,88,91,92,93,144,170,80,81,146,134,91,132,83,88,100,82,143,90,91,90,150,129,")
	data = append(data, "77,84,84,107,88,87,92,119,144,95,150,118,90,88,154,146,84,83,80,134,138,98,85,88,85,147,119,85,86,85,90,125,100,88,91,86,86,82,84,165,81,89,84,156,153,93,150,143,82,107,106,151,148,88,149,113,87,163,")
	data = append(data, "77,84,84,87,146,145,152,93,83,82,87,113,150,88,91,128,97,102,81,81,171,178,98,181,178,103,88,112,86,85,124,94,90,88,125,96,115,96,90,80,90,128,98,156,116,87,92,85,86,150,153,101,86,84,87,82,83,82,142,")
	data = append(data, "77,84,84,87,146,133,170,145,83,82,87,88,158,84,118,86,87,149,144,152,103,84,148,90,84,84,89,81,86,119,98,169,106,88,131,92,93,92,80,80,146,88,88,91,88,87,92,85,86,85,90,91,90,156,154,")
	data = append(data, "77,84,84,87,84,83,91,150,89,101,87,88,87,85,131,146,87,86,146,87,78,98,110,88,156,147,182,175,82,95,90,125,93,146,87,82,83,82,148,84,81,88,88,91,88,87,92,85,86,85,90,91,90,88,91,86,155,94,")
	data = append(data, "77,84,84,87,84,83,88,143,86,95,150,88,87,85,88,154,83,86,84,146,141,85,98,87,97,83,88,81,82,85,134,87,101,91,87,82,83,96,84,84,81,88,88,91,88,87,92,85,86,85,90,91,90,88,91,86,87,86,84,147,")
	data = append(data, "77,84,84,87,84,83,88,81,82,143,90,154,150,85,88,83,155,146,84,84,91,104,178,87,84,83,88,81,82,85,90,91,168,91,87,82,83,110,84,84,81,88,88,91,88,87,92,85,86,85,90,91,90,88,91,86,87,130,142,142,139,")
	data = append(data, "77,84,84,87,84,83,88,81,82,81,86,149,170,153,101,146,177,82,84,84,143,84,84,87,84,83,88,81,82,145,124,172,155,97,87,82,83,146,84,84,81,88,88,91,88,87,92,85,86,85,90,91,134,146,101,108,83,82,81,81,77,146,112,")
	data = append(data, "77,84,84,87,84,83,88,81,82,81,86,87,86,84,149,144,112,102,84,128,139,84,84,87,84,83,102,93,94,93,100,87,99,146,101,82,83,82,111,84,81,88,88,91,88,87,92,85,86,143,117,113,100,84,88,83,83,108,81,81,171,84,96,99,153,150,102,95,")
	data = append(data, "77,84,84,87,84,83,88,81,82,81,86,87,86,84,87,82,83,82,80,80,91,98,174,91,88,87,92,85,86,85,90,91,158,96,87,82,83,82,144,84,81,88,88,91,88,131,150,95,96,145,87,157,86,110,88,152,83,176,89,124,167,153,88,91,88,131,150,143,144,143,117,160,")
	data = append(data, "77,84,84,87,84,83,88,81,82,81,86,87,86,84,87,82,83,82,165,148,81,88,88,91,88,87,92,85,86,85,90,91,90,88,159,82,83,82,80,144,108,88,88,125,84,111,89,82,89,109,87,117,112,178,96,167,152,86,84,128,87,146,98,101,110,84,89,82,83,82,86,129,")
	data = append(data, "77,84,84,87,84,83,88,81,82,81,86,87,86,84,87,82,97,125,84,84,81,88,88,91,88,87,92,85,86,85,90,91,90,88,91,94,83,82,80,80,77,132,98,117,153,90,89,151,82,82,180,87,155,152,91,86,131,144,94,106,78,85,85,88,85,84,89,82,146,95,154,149,")
	data = append(data, "77,84,84,87,84,83,88,81,82,81,86,87,86,84,87,172,87,86,84,84,81,88,88,91,88,87,92,85,86,85,90,91,90,88,91,86,83,82,80,80,77,84,104,88,85,177,142,150,86,85,90,135,150,110,88,83,84,83,81,81,78,85,85,88,91,97,157,129,89,")
	data = append(data, "77,84,84,87,84,83,88,81,82,81,86,87,86,84,100,92,145,96,80,94,77,146,115,118,88,87,92,85,86,85,90,91,90,88,91,86,83,82,80,80,77,84,84,181,84,147,178,85,86,85,124,149,100,148,88,83,84,83,81,81,78,85,85,88,91,95,156,129,144,")
	data = append(data, "77,84,84,87,84,83,88,81,82,81,86,87,100,146,87,91,109,83,81,152,77,157,148,88,97,83,96,143,113,85,90,91,90,88,91,152,83,82,80,80,77,84,84,87,96,87,92,129,144,107,87,87,95,98,87,146,90,83,81,81,78,148,96,155,88,121,95,")
	data = append(data, "77,84,84,87,84,83,88,95,167,150,155,156,171,98,129,146,96,144,94,81,149,84,112,100,84,83,182,145,96,95,86,118,117,88,91,89,83,82,80,80,77,84,84,99,88,87,88,107,83,82,150,156,90,88,91,86,152,82,144,92,146,88,122,94,")
	data = append(data, "77,84,84,87,153,87,92,85,86,85,90,91,90,88,91,86,87,86,149,149,150,119,84,149,93,114,114,94,110,107,91,94,112,146,118,82,83,82,80,80,77,84,84,118,88,153,88,82,83,82,86,118,90,88,91,86,145,94,118,142,")
	data = append(data, "77,84,84,135,146,97,102,95,96,145,150,151,148,146,118,86,87,86,84,84,81,88,88,91,153,168,132,143,96,82,100,101,150,84,101,82,83,82,80,80,77,84,84,87,146,135,173,95,145,95,98,151,150,146,149,")
	data = append(data, "77,84,84,91,84,84,89,82,83,82,87,88,87,85,88,108,109,96,142,142,108,115,88,91,88,87,92,85,151,123,180,149,115,98,87,82,83,82,80,80,77,84,84,87,84,83,152,145,146,145,")
	data = append(data, "77,84,84,118,88,97,152,82,83,82,87,88,87,85,88,83,84,83,81,81,78,110,98,149,115,114,92,85,86,85,90,154,93,98,")
	data = append(data, "77,84,84,87,146,114,154,95,146,82,87,88,87,85,88,83,84,83,81,81,78,85,85,88,85,113,102,143,113,85,90,91,90,151,")
	data = append(data, "77,84,84,87,84,83,150,112,86,148,150,94,87,85,88,83,84,83,81,81,78,85,85,94,96,152,157,147,94,81,148,91,90,88,162,")
	data = append(data, "77,84,84,87,84,83,88,81,82,143,117,91,155,98,151,89,84,83,81,81,78,85,85,87,88,87,92,85,86,81,113,101,117,88,91,")
	data = append(data, "77,84,84,87,84,83,88,81,82,81,86,87,148,115,91,151,152,94,144,144,84,91,85,101,146,114,92,129,146,154,87,88,86,103,135,")
	data = append(data, "77,84,84,87,84,83,88,81,82,81,86,87,86,84,87,82,145,144,111,111,81,150,153,154,96,95,102,93,149,150,155,156,134,146,")
	data = append(data, "77,84,84,87,84,83,88,81,82,81,86,87,86,84,87,82,83,82,80,80,77,84,84,151,146,145,150,143,144,143,")
	data = append(data, "77,")
	data = append(data, "122,153,166,169,173,83,123,153,164,154,169,171,163,149,170,82,148,160,148,80,154,149,173,87,126,148,166,166,147,163,175,87,103,167,171,82,149,151,80,155,150,162,152,87,168,162,88,170,161,166,168,87,")
	return data
}
