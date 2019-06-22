<template>
  <div>
    <!-- 
      Authモジュールのコマンド
      $auth.loggedIn ログインしているか否か
      $auth.user     ユーザー情報
      $auth.logout   ログアウト
     -->
    <div style="height: 28px; padding: 8px; border-bottom: solid 2px #000;">
      <span>{{ $auth.loggedIn ? $auth.user.Name : 'Guest' }}</span>
      <button 
        v-if="$auth.loggedIn"
        style="float:right"
        v-on:click="$auth.logout()"
      >
        ログアウト
      </button>
    </div>
    <nuxt />
  </div>
</template>