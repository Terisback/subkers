import './geometry.css'
import anime from "animejs"

const i = Math.floor(Math.random() * 3)
const ids = ['ellipse', 'rectangle', 'triangle']

window.onload = () => {
    const bg = document.getElementById('geometry')
    const frag = document.createDocumentFragment()
    for (let x = 0; x < 20; x++) {
        const row = document.createElement('div')
        for (let y = 0; y < 20; y++) {
            const figure = ids[(i + x + y) % 3];
            const d = document.createElement('div')
            d.classList.add('el', figure)
            const e = document.createElementNS('http://www.w3.org/2000/svg', 'svg')
            e.setAttribute('width', '22')
            e.setAttribute('height', '22')
            e.setAttribute('version', '2.0')
            const u = document.createElementNS('http://www.w3.org/2000/svg', 'use')
            u.setAttribute('href', '#' + figure)
            e.appendChild(u)
            d.appendChild(e)
            row.appendChild(d)
        }
        frag.appendChild(row)
    }
    bg.appendChild(frag)

    anime({
        targets: '.el svg',
        rotate: [
            { value: 180, easing: 'easeOutSine', duration: 1300 },
            { value: 360, easing: 'easeInOutQuad', duration: 700 }
        ],
        loop: true,
        delay: anime.stagger(100, { grid: [20, 20], from: 'center', direction: 'reverse' }),
    })

    anime({
        targets: ['.el.rectangle', '.el.ellipse'],
        scale: [
            { value: .6, easing: 'easeOutSine', duration: 2000 },
            { value: .8, easing: 'easeInOutQuad', duration: 2000 }
        ],
        loop: true,
        direction: 'alternate',
        delay: anime.stagger(100, { grid: [20, 20], from: 'center' }),
    })

    anime({
        targets: '.el.triangle',
        scale: [
            { value: .7, easing: 'easeOutSine', duration: 2000 },
            { value: 1, easing: 'easeInOutQuad', duration: 2000 }
        ],
        loop: true,
        delay: anime.stagger(100, { grid: [20, 20], from: 'center' }),
    })
}