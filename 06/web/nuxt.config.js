export default {
  mode: 'universal',

  /*
  ** Headers of the page
  */
  head: {
    
    title: process.env.npm_package_name || '',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: process.env.npm_package_description || '' }
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }
    ]
  },

  /*
  ** Customize the progress-bar color
  */
  loading: { color: '#fff' },

  /*
  ** Global CSS
  */
  css: [
  ],

  /*
  ** Plugins to load before mounting the App
  */
  plugins: [
  ],

  /*
  ** Nuxt.js modules
  */
  modules: [
    '@nuxtjs/auth', // Authモジュール https://auth.nuxtjs.org/
    '@nuxtjs/axios',
  ],
  // Authモジュールの設定です。
  auth: {
    redirect: {
      logout: '/login', // ログアウト後のリダイレクト先を指定しています。
    },
    strategies: {
      local: {
        // ログイン、ログアウト、ユーザー情報取得の際のAPIのエンドポイントを指定します。
        // propatyNameは、レスポンスボディのJSONのうち、どのプロパティからデータを取得するかを指定しています。
        endpoints: {
          login: { url: '/user', method: 'put', propertyName: 'token' },
          // ログアウト時にはAPIへのリクエストが不要なので、falseを指定します。
          logout: false,
          // Authモジュールは、ログイン後にユーザー情報を取得する仕様になっています。
          // userのpropatyNameがfalseなのは、レスポンスボディのJSONがそのままユーザー情報になるという意味です。
          // this.$auth.fetchUser()でもユーザー情報を取得することができます。
          user: { url: '/auth/user', method: 'get', propertyName: false },
        },
        tokenType: 'bearer',
      },
    },
  },
  axios: {
    baseURL: 'http://api:5000',
    browserBaseURL: 'http://localhost:15000',
  },
  
  /*
  ** Build configuration
  */
  build: {
    /*
    ** You can extend webpack config here
    */
    extend(config, ctx) {
    }
  },
  // Docker for Windowsでは、ホットリロードを有効にするために、下記の設定を追加します。
  watchers: {
    webpack: {
      poll: true
    }
  },
}
