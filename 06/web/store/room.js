export const state = () => ({
  room: {},
})

export const getters = {
  room(state) {
    return state.room
  },
  comments(state) {
    return state.room.Comments || []
  },
}

export const mutations = {
  SET_ROOM(state, newRoom) {
    state.room = newRoom
  },
  UNSHIFT_COMMENT(state, newComment) {
    if (!state.room.Comments) state.room.Comments = []
    state.room.Comments.unshift(newComment)
  },
  REMOVE_COMMENT(state, index) {
    state.room.Comments.splice(index, 1)
  },
}

export const actions = {
  async getRoom({ commit }, roomId) {
    try {
      const { data } = await this.$axios.get(`/auth/room/${roomId}`)
      commit('SET_ROOM', data)
    } catch (err) {
      console.error(err)
      return err
    }
  },
  async createComment({ getters, commit }, body) {
    try {
      const { data } = await this.$axios.post(`/auth/room/${getters.room.ID}/comment`, { body })
      commit('UNSHIFT_COMMENT', data)
    } catch (err) {
      console.error(err)
      return err
    }
  },
  async removeComment({ getters, commit }, index) {
    try {
      await this.$axios.delete(`/auth/comment/${getters.comments[index].ID}`)
      commit('REMOVE_COMMENT', index)
    } catch (err) {
      console.error(err)
      return err
    }
  },
}