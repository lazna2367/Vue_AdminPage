import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
    state:{
        items:[]
    },
    getters: {

    },
    mutations: {

    },
    actions:{
        updateClick(){
            console.log('도니?2');
            const baseURI = 'http://172.30.1.46:8091/depositList'
            this.$axios.post(`${baseURI}`)
            .then((result) => {          
               this.state.items = result.data
            }) 
        }
    }
})