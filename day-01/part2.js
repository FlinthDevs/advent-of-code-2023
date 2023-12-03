const fs = require('fs')

const input = fs.readFileSync('./input', 'utf8')
const lines = input.split('\n')

const letteredNumbers = {
	"one": 1,
	"two": 2,
	"three": 3,
	"four": 4,
	"five": 5,
	"six": 6,
	"seven": 7,
	"eight": 8,
	"nine": 9,
}

const abbrs = Object.keys(letteredNumbers).reduce(
	(prev, val) => {
		prev[val.replace(/.$/, "")] = val
		return prev
		}
	, {}
	)

const lettersRegex = RegExp(`${Object.keys(letteredNumbers).reduce((prev,val) => `${prev}|${val.replace(/.$/, (m) => `(?=${m})`)}`, "").replace(/^\|/, "")}`, "g")

const numbersRegex = RegExp("[0-9]", "g")
let sum = 0

for (let i = 0; i < lines.length; i++) {
	if (lines[i].length === 0) { continue }
	console.log(`Line ${i}: `, lines[i])

	const matches = [...lines[i].matchAll(numbersRegex)]

	let firstNumber = lastNumber = {index: -1}

	if (matches.length > 0) {
		firstNumber = matches[0]
		lastNumber = matches[matches.length - 1]
	}

	const letterMatches = [...lines[i].matchAll(lettersRegex)]
	let firstLetteredNumber = { index: -1 }
	let lastLetteredNumber = { index: -1 }

	if (letterMatches.length > 0) {
		firstLetteredNumber = letterMatches[0]
		lastLetteredNumber = letterMatches[letterMatches.length - 1]
	}

	if (firstLetteredNumber.index === 0 || (firstLetteredNumber.index >= 0 && firstLetteredNumber.index < firstNumber.index)) {
		firstNumber = letteredNumbers[abbrs[firstLetteredNumber[0]]] + ""
	}

	if (lastLetteredNumber.index > lastNumber.index) {
		lastNumber = letteredNumbers[abbrs[lastLetteredNumber[0]]] + ""
	}

	const toAdd = firstNumber + lastNumber
	console.log("  Added: ", toAdd)
	sum += parseInt(toAdd)
}

console.log("Result: ", sum)
