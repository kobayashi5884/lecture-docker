<template>
<!-- ここにはHTMLを書きますが、Vue.jsならではの便利な書き方ができます。 -->
<!-- templateの直下には１つの要素（ここではdiv）しか定義できませんので、注意してください。 -->
  <div>
    <div>{{ text }}</div>
    <!-- アプリケーション内でのページ遷移（SPAによるページ遷移）には、nuxt-linkとtoを使います。 -->
    <nuxt-link to="/room">掲示板</nuxt-link>
  </div>
</template>

<script>
// ここにはJavaScriptを書きますが、Vue.jsのルールで書きます。
export default {
  data() {
    return {
      text: 'Hello World',
    }
  },
}
</script>

<style>
/* ここには通常のCSSを書きます。 */
</style>
