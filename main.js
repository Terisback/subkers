import './style.css'

import swal from 'sweetalert';
import { parse } from 'ass-compiler';
import srtParser from 'srt-parser-2';
import { saveAs } from "file-saver";

const fileTypes = ["srt", "ssa", "ass"]
const newlineRegexp = /[\\N\n]/g

function checkFileType(file) {
  const ext = file.name.split('.').pop()
  console.log(ext)
  for (var i = 0; i < fileTypes.length; i++) {
    if (ext == fileTypes[i]) {
      return true
    }
  }
  return false
}

function convertAssTime(time) {
  const min = Math.floor(time / 60)
  const sec = Math.floor(time % 60)
  const milli = Math.floor((time - Math.floor(time)) * 1000)
  return `${min}:${sec}.${milli}`
}

function convertSrtTime(srtTime) {
  const hour = +srtTime.slice(0, 1)
  const min = +srtTime.slice(3, 4)
  const sec = srtTime.slice(6, 7)
  const milli = srtTime.slice(9, 11)
  return `${hour * 60 + min}:${sec}.${milli}`
}

function onChange(event) {
  let file = event.target.files[0]

  if (!checkFileType(file)) {
    swal('Wrong file extension!', 'Supported file extensions: .srt, .ssa, .ass', 'error')
    return
  }

  let reader = new FileReader()
  reader.onload = async function (e) {
    const filename = file.name.split('.')
    const extension = filename.pop()
    console.log(extension)
    let markers = 'Name\tStart\tEnd\tTime Format\tType\tDescription\n'

    switch (extension) {
      case 'srt':
        const parser = new srtParser()
        const subtitles = parser.fromSrt(e.target.result)

        subtitles.forEach((subtitle, i) => {
          markers += `${subtitle.text.replace(newlineRegexp, ' ')}\t${convertSrtTime(subtitle.startTime)}\t${convertSrtTime(subtitle.endTime)}\tdecimal\tCue\t\n`
        })
        break

      case 'ass':
      case 'ssa':
        const parsed = parse(e.target.result)
        parsed.events.dialogue.forEach(subtitle => {
          markers += `${subtitle.Text.combined.replace(newlineRegexp, ' ')}\t${convertAssTime(subtitle.Start)}\t${convertAssTime(subtitle.End)}\tdecimal\tCue\t\n`
        })
        break
      default:
        swal('Wrong file extension!', 'Supported file extensions: .srt, .ssa, .ass', 'error')
        break
    }

    var blob = new Blob([markers], {
      type: 'text/plain;charset=utf-8',
    })
    saveAs(blob, filename + '.csv')
  };

  reader.readAsText(file);
}

document.getElementById('file-input').addEventListener('change', onChange)