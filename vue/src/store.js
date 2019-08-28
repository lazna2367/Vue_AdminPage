import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
    state:{
        items:[],
        web3: {
            isInjected: false,
            web3Instance: null,
            networkId: null,
            coinbase: null,
            balance: null,
            error: null
        },
        contractInstance: null,
    },
    getters: {

    },
    mutations: {

    },
    actions:{
        updateClick(){            
            const baseURI = 'http://172.30.1.46:8091/depositList'
            this.$axios.post(`${baseURI}`)
            .then((result) => {          
               this.state.items = result.data
            }) 
        }
    }
})