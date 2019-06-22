<template>
  <div>
    <h2>スレッド一覧</h2>

    <div style="margin-bottom: 16px;">
      <input 
        v-model="title" 
        type="text"
        placeholder="タイトル"
      >
      <button v-on:click="create">スレッド作成</button>
    </div>

    <ul>
      <li 
        v-for="(room, i) in rooms" 
        v-bind:key="room.ID"
      >
        <nuxt-link v-bind:to="`/room/${room.ID}`">{{ room.Title }}</nuxt-link>
        <button 
          v-if="$auth.user.ID == room.UserID"
          v-on:click="remove(i)"
        >
          スレッド削除
        </button>
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  middleware: 'auth',
  data() {
    return {
      title : '',
    }
  },
  // ストアのgetterを参照するときは、computedを使います。
  // gettersのプロパティ名は'ストアのファイル名（拡張子不要）/getter名'です。
  // ファイル名がindex.jsの場合には、getter名のみで参照できます。
	computed: {
		rooms() {
			return this.$store.getters['rooms/rooms']
		},
	},
  // APIから取得したデータをdataではなくストアに保存するときは、asyncDataではなく、fetchを使います。
  async fetch({ store, error }) {
    // ストアのactionを呼び出すときは、dispatchを使います。
    // 引数は'ストアのファイル名（拡張子不要）/action名'です。
    // ファイル名がindex.jsの場合には、action名のみで呼び出せます。
    const err = await store.dispatch('rooms/getRooms')
    if (err) {
			error({
        statusCode: err.response.status,
        message:    err.response.statusText,
      })
    }
  },
  methods: {
    async create() {
      // fetchとmethodsでactionの呼び出し方が少し違いますので、注意してください。
      const err = await this.$store.dispatch('rooms/createRoom', this.title)
      if (err) {
        alert('スレッドの作成に失敗しました')
      }
    },
    async remove(i) {
      const err = await this.$store.dispatch(`rooms/removeRoom`, i)
      if (err) {
        alert('スレッドの削除に失敗しました')  
      }
    },
  },
}
</script>
