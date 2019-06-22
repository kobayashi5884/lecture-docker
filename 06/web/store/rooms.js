// stateには生のデータを保存します。
export const state = () => ({
  rooms: [],
})

// gettersにはコンポーネントから参照する生のデータまたはそれを加工したデータを定義します。
export const getters = {
  rooms(state) {
    return state.rooms
  },
}

// mutationsにはstateを変更する関数を定義します。
// mutations以外でstateを変更すると、エラーになります。
export const mutations = {
  SET_ROOMS(state, newRooms) {
    state.rooms = newRooms
  },
  PUSH_ROOM(state, newRoom) {
    state.rooms.push(newRoom)
  },
  REMOVE_ROOM(state, index) {
    state.rooms.splice(index, 1)
  },
}

// actionsにはAPIとの通信その他の処理を行う関数を定義します。
export const actions = {
  async getRooms({ commit }) {
    try {
      const { data } = await this.$axios.get('/auth/room')
      commit('SET_ROOMS', data)
    } catch (err) {
      console.error(err)
      return err
    }
  },
  async createRoom({ commit }, title) {
    try {
      const { data } = await this.$axios.post('/auth/room', { title })
      commit('PUSH_ROOM', data)
    } catch (err) {
      console.error(err)
      return err
    }
  },
  async removeRoom({ getters, commit }, index) {
    try {
      await this.$axios.delete(`/auth/room/${getters.rooms[index].ID}`)
      commit('REMOVE_ROOM', index)
    } catch (err) {
      console.error(err)
      return err
    }
  },
}