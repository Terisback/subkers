
interface Line {
    Start: number
    End: number
    Text: string
}

const lineRegex = /(\d+)\n(\d{2}:\d{2}:\d{2}[,.]\d{3}) --> (\d{2}:\d{2}:\d{2}[,.]\d{3})/g;

export function parseSRT(data: string): Line[] {
    let lines = data.replace(/\r/g, '').split(lineRegex);
    lines.shift();
    return lines
        .reduce((resultArray: string[][], item, index) => {
            const chunkIndex = Math.floor(index / 4)
            if (!resultArray[chunkIndex]) resultArray[chunkIndex] = [] // start a new chunk
            resultArray[chunkIndex].push(item)
            return resultArray
        }, [])
        .map(chunk => {
            const line: Line = {
                Start: timeMs(chunk[1].trim()),
                End: timeMs(chunk[2].trim()),
                Text: chunk[3].trim(),
            }
            console.log(line)
            return line
        })
}

const timeRegex = /(\d+):(\d{2}):(\d{2})[,.](\d{3})/;

function timeMs(val: string): number {
    let parts = timeRegex.exec(val);


    let numbers: number[] = parts!.map(part => {
        const num = parseInt(part!, 10)
        if (isNaN(num)) return 0
        else return num
    })

    console.log(parts, numbers)


    // hours + minutes + seconds + ms
    return numbers[1] * 3600000 + numbers[2] * 60000 + numbers[3] * 1000 + numbers[4];
}