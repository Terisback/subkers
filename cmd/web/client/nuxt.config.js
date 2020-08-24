
export default {
  /*
  ** Nuxt rendering mode
  ** See https://nuxtjs.org/api/configuration-mode
  */
  mode: 'spa',
  /*
  ** Nuxt target
  ** See https://nuxtjs.org/api/configuration-target
  */
  target: 'static',
  /*
  ** Headers of the page
  ** See https://nuxtjs.org/api/configuration-head
  */
  head: {
    htmlAttrs: {
      lang: 'en'
    },
    title: "Subkers",
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { name: 'apple-mobile-web-app-title', content: "Subkers" },
      { name: 'application-name', content: "Subkers" },
      { name: 'msapplication-TileColor', content: "#2b5797" },
      { name: 'theme-color', content: '#ffffff'},
      { hid: 'description', name: 'description', content: "Tool for converting subtitles into markers for Adobe Audition" }
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' },
      { rel: "icon", type: "image/png", href: "/favicon-32x32.png"},
      { rel: "icon", type: "image/png", href: "/favicon-16x16.png"},
      { rel: "manifest", href: "/site.webmanifest"},
      { rel: "mask-icon", href: "/safari-pinned-tab.svg", color: "#5bbad5"},
      { rel: "apple-touch-icon", sizes: "180x180", href: "/apple-touch-icon.png"}
    ]
  },
  /*
  ** Global CSS
  */
  css: [
  ],
  /*
  ** Plugins to load before mounting the App
  ** https://nuxtjs.org/guide/plugins
  */
  plugins: [
  ],
  /*
  ** Auto import components
  ** See https://nuxtjs.org/api/configuration-components
  */
  components: true,
  /*
  ** Nuxt.js dev-modules
  */
  buildModules: [
  ],
  /*
  ** Nuxt.js modules
  */
  modules: [
  ],
  /*
  ** Build configuration
  ** See https://nuxtjs.org/api/configuration-build/
  */
  build: {
  }
}
