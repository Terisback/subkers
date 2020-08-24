/* eslint-disable */

export default ({ app }) => {
    /*
     ** Будет работать только на стороне клиента и только когда режим разработки будет "production"
     */
    if (process.env.NODE_ENV !== 'production')
        return /*
       ** Подключаем скрипт Google Analytics
       */
            ; (function (i, s, o, g, r, a, m) {
                i['GoogleAnalyticsObject'] = r
                    ; (i[r] =
                        i[r] ||
                        function () {
                            ; (i[r].q = i[r].q || []).push(arguments)
                        }),
                        (i[r].l = 1 * new Date())
                    ; (a = s.createElement(o)), (m = s.getElementsByTagName(o)[0])
                a.async = 1
                a.src = g
                m.parentNode.insertBefore(a, m)
            })(
                window,
                document,
                'script',
                'https://www.google-analytics.com/analytics.js',
                'ga'
            )
    /*
     ** Устанавливаем текущую страницу
     */
    ga('create', 'UA-111363587-2', 'auto')
    /*
     ** Вызывается каждый раз после смены роута (при инициализации тоже)
     */
    app.router.afterEach((to, from) => {
        /*
         ** We tell Google Analytics to add a `pageview`
         */
        ga('set', 'page', to.fullPath)
        ga('send', 'pageview')
    })
}