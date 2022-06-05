import './style.css'

import swal from 'sweetalert'
import { parse as parseASS } from 'ass-compiler'
import { saveAs } from "file-saver"
import { fromVtt } from "subtitles-parser-vtt";

const fileTypes = ["srt", "ssa", "ass"]
const newlineRegexp = /[\\N\n]/g

function checkFileType(file: File) {
  const ext = file.name.split('.').pop()

  for (const filetype of fileTypes) {
    console.log(filetype)
    if (ext == filetype) return true
  }

  return false
}

function milliToTime(time: number) {
  const min = Math.floor(time / 60)
  const sec = Math.floor(time % 60)
  const milli = Math.floor((time - Math.floor(time)) * 1000)
  return `${min}:${sec}.${milli}`
}

type Subtitle = {
  // Cue start, in milliseconds
  Start: number
  // Cue duration, in milliseconds
  Duration: number
  Text: string
}

function onChange(event: Event) {
  let input = (event.target! as HTMLInputElement)
  let files = Array.from(input.files!)

  for (const file of files) {
    if (!checkFileType(file)) {
      swal('Wrong file extension!', 'Supported file extensions: .srt, .ssa, .ass', 'error')
      return
    }
  }

  // Processing files
  for (const file of files) {
    let reader = new FileReader()
    reader.onload = async function (e) {
      const filename = file.name.split('.')
      const extension = filename.pop()

      if (!e.target) {
        swal('Failed to read!', 'Failed to read file')
        return
      }

      let subtitles: Subtitle[] = []
      switch (extension) {
        case 'srt':
          subtitles = fromVtt(e.target!.result!.toString())
            .map(subtitle => {
              const sub: Subtitle = {
                Start: subtitle.startTime,
                Duration: subtitle.endTime - subtitle.startTime,
                Text: subtitle.text.replace(newlineRegexp, ' ')
              }
              return sub
            })
          break

        case 'ass':
        case 'ssa':
          subtitles = parseASS(e.target!.result!.toString()).events.dialogue
            .map(entry => {
              const sub: Subtitle = {
                Start: entry.Start,
                Duration: entry.End - entry.Start,
                Text: entry.Text.combined.replace(newlineRegexp, ' ')
              }
              return sub
            })
          break

        default:
          swal('Wrong file extension!', 'Supported file extensions: .srt, .ssa, .ass', 'error')
          break
      }

      // Saving to CSV
      let markers = subtitles.map(subtitle => `${subtitle.Text.replaceAll('"', '""')}\t${milliToTime(subtitle.Start)}\t${milliToTime(subtitle.Duration)}\tdecimal\tCue\t`)
      markers.unshift('Name\tStart\tEnd\tTime Format\tType\tDescription\n')
      var blob = new Blob([markers.join('\n')], {
        type: 'text/plain;charset=utf-8',
      })
      saveAs(blob, filename + '.csv')
    }
    reader.readAsText(file)
  }
}

document.getElementById('file-input')!.addEventListener('change', onChange)