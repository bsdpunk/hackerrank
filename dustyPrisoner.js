process.stdin.resume();
process.stdin.setEncoding('ascii');

var input_stdin = "";
var input_stdin_array = "";
var input_currentline = 0;

process.stdin.on('data', function (data) {
    input_stdin += data;
});

process.stdin.on('end', function () {
    input_stdin_array = input_stdin.split("\n");
    main();    
});

function readLine() {
    return input_stdin_array[input_currentline++];
}

/////////////// ignore above this line ////////////////////

function saveThePrisoner(n, m, s){
	psion = m + s
	dbl = n * 2
	//var final int64 := psion
	if (s > dbl) {
		//pset.sweets
		if (psion % n == psion) {
			//			fmt.Println("1")
			return psion -1 
		}
		if (((psion -1 ) % n) == 0) {
			//			fmt.Println("2")
			return n
		}
		//if (psion % pset.total) < 0 {
		//			fmt.Println("3")
		//	return psion, nil
		//}

		//	}
		//	fmt.Println("4")
		//fmt.Println(psion)
		return ((psion - 1) % n)
	} else if (psion > n) {

		psion = -1*(n - psion) - 1
		if (psion -1 == n) {
			return 1
		}
		if (psion == 0) {
			return n
		}

		if (psion > n) {
                        console.log(psion + " - " + n)
			return psion - n
		}
		return psion
	} else {
		return psion - 1 
	}
}


function main() {
    var t = parseInt(readLine());
    for(var a0 = 0; a0 < t; a0++){
        var n_temp = readLine().split(' ');
        var n = parseInt(n_temp[0]);
        var m = parseInt(n_temp[1]);
        var s = parseInt(n_temp[2]);
        var result = saveThePrisoner(n, m, s);
        process.stdout.write(""+result+"\n");
    }

}
