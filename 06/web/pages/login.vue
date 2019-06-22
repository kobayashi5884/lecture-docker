<template>
  <div>
    <h2>ログインまたは新規登録</h2>
    <input 
      v-model="name" 
      type="text"
      placeholder="名前"
    >
    <input 
      v-model="password" 
      type="text"
      placeholder="パスワード"
    >
    <button v-on:click="login">ログイン</button>
    <button v-on:click="signup">新規登録</button>
  </div>
</template>

<script>
export default {
  // 既にログイン状態であれば、ホームページにリダイレクトするというAuthモジュールのコマンドです。
  auth: 'guest',
  data() {
    return {
      name: '',
      password: '',
    }
  },
  methods: {
    async signup() {
      try {
        const { data } = await this.$axios.post('/user', {
          name:     this.name.trim(),
          password: this.password,
        })
        // AuthモジュールにJWTをセットします。
        this.$auth.setToken('local', `bearer ${data.token}`)
      } catch (err) {
        console.error(err)
        // レスポンスボディからエラーメッセージを取得し、表示します。
        const msg = err.response.data.trim() || '新規登録に失敗しました。'
        alert(msg)
        return
      }

      // AuthモジュールにセットしたJWTでユーザー情報を取得し、ログイン状態にします。
      this.$auth.setStrategy('local')
      await this.$auth.fetchUser()

      // ホームページにリダイレクトします。
      this.$router.push('/')
    },
    async login() {
      // Authモジュールのヘルパー関数でログインします。
      try {
        await this.$auth.loginWith('local', {
          data: {
            name:     this.name.trim(),
            password: this.password,
          }
        })
      } catch (err) {
        console.error(err)
        const msg = err.response.data.trim() || 'ログインに失敗しました。'
        alert(msg)
      }
    },
  },
}
</script>