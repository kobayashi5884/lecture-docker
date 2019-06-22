<template>
  <div>
    <!-- 
      時々、
      'The client-side rendered virtual DOM tree is not matching server-rendered content...'
      というエラーがコンソールに表示される場合があります。
      どうやらSSRでうまくページを描画できないということらしいので、
      原因不明の場合には、最終手段として、no-ssrタグで囲むと、エラーが表示されなくなります。
     -->
    <no-ssr>
      <div 
        v-for="(comment, i) in comments" 
        v-bind:key="comment.ID"
        style="margin-bottom: 8px;"
      >
        <div>
          <div style="display: inline-block; font-weight: bold;">
            {{ comment.User.Name }}
          </div>
          <div style="display: inline-block; margin-left: 8px">
            {{ comment.CreatedAt }}
          </div>
        </div>
        {{ comment.Body }}
        <!-- 
          $emitでインポート先にデータを渡します。
          第1引数がイベント名で、第二引数がインポート先の関数に渡される引数です。
        -->
        <button 
          v-show="$auth.user.ID == comment.UserID"
          v-on:click="$emit('remove', i)"
        >
          削除
        </button>
      </div>
    </no-ssr>
  </div>
</template>

<script>
export default {
  // インポート先から属性として受け取ったデータは、propsで参照します。
  // props内でデータのバリデーションをしたり、デフォルトを指定したりすることもできます。
  props: {
		comments: {
			type: Array,
			default: [],
		},
	},
}
</script>