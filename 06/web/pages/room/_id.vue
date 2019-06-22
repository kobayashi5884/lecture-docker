<template>
  <div>
    <h2>{{ room.Title }}</h2>

    <div style="margin-bottom: 16px;">
      <input 
        v-model="comment" 
        type="text"
        placeholder="コメント"
      >
      <button v-on:click="create">送信</button>
    </div>

    <!-- 
      インポートしたコンポーネントを下記のようにカスタムタグとして利用できます。
      属性の形式でコンポーネントにデータを渡し、
      イベントの形式でコンポーネントからデータを受け取ります。
      なお、今回の例では、属性やイベントの形式ではなく、
      ストアを利用してデータの受け渡しもできますので、考えてみてください。
     -->
    <Comments 
      v-bind:comments="comments"
      v-on:remove="remove"
    />
  </div>
</template>

<script>
// コンポーネントのファイルパスを指定してインポートします。
import Comments from "@/components/Comments"

export default {
  middleware: 'auth',
  // インポートするだけではなく、下記のようにコンポーネントを指定する必要があります。
  components: { Comments },
  data() {
    return {
      comment : '',
    }
  },
	computed: {
		room() {
			return this.$store.getters['room/room']
		},
		comments() {
			return this.$store.getters['room/comments']
		},
  },
  // URLのパラメーターにはparamsでアクセスできます。
  // ファイル名を_id.vueにしているので、params.idでURLの/room/に続く数値を参照できます。
  async fetch({ store, params, error }) {
    const err = await store.dispatch('room/getRoom', params.id)
    if (err) {
			error({
        statusCode: err.response.status,
        message:    err.response.statusText,
      })
    }
  },
  methods: {
    async create() {
      const err = await this.$store.dispatch('room/createComment', this.comment)
      if (err) {
        alert('コメントの作成に失敗しました')
      }
    },
    async remove(i) {
      const err = await this.$store.dispatch(`room/removeComment`, i)
      if (err) {
        alert('コメントの削除に失敗しました')  
      }
    },
  },
}
</script>

