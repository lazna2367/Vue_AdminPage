<template>
  <b-card>    
    <div slot="header" v-html="caption"></div>
    <b-table :dark="dark" :hover="hover" :striped="striped" :bordered="bordered" :small="small" :fixed="fixed" responsive="sm" :items="items" :fields="captions" :current-page="currentPage" :per-page="perPage">            
      <!-- <template slot="header-name" slot-scope="data">
          <b-form-checkbox></b-form-checkbox>
      </template> -->
      <!-- <template slot="status" slot-scope="data">        
        <b-badge :variant="getBadge(data.item.status)">{{data.item.status}}</b-badge>                
      </template> -->          
      <template slot="Complete" slot-scope="data">  
        <!-- pending -->
        <b-dropdown v-if="data.item.Complete == 'not'" size="sm" id="ddown_secondary" variant="secondary" class="m-1" text="미확인">        
          <b-dropdown-item @click="onClick(data.item.Id,'confirmed','not')">미승인</b-dropdown-item>
          <b-dropdown-item @click="onClick(data.item.Id,'confirmed','pending')">보류</b-dropdown-item>
          <b-dropdown-item @click="onCompleted(
            data.item.Id,
            'confirmed',
            'completed',
            data.item.Wallet.Wallet_address,
            data.item.Base_unit,
            data.item.Base_amount,
            data.item.Target_unit,
            data.item.Target_amount,
            data.item.Rate_by_base,
            data.item.Payout_account.Transaction
            )">승인</b-dropdown-item>
        </b-dropdown>        
        <b-dropdown v-if="data.item.Complete == 'pending'" size="sm" id="ddown_warning" variant="warning" class="m-1" text="대기중">        
          <b-dropdown-item @click="onClick(data.item.Id,'confirmed','not')">미승인</b-dropdown-item>
          <b-dropdown-item @click="onClick(data.item.Id,'confirmed','pending')">보류</b-dropdown-item>
          <b-dropdown-item @click="onCompleted(
            data.item.Id,
            'confirmed',
            'completed',
            data.item.Wallet.Wallet_address,
            data.item.Base_unit,
            data.item.Base_amount,
            data.item.Target_unit,
            data.item.Target_amount,
            data.item.Rate_by_base,
            data.item.Payout_account.Transaction
            )">승인</b-dropdown-item>
        </b-dropdown>
      </template>
    </b-table>
    
    <nav>
      <b-pagination :total-rows="totalRows" :per-page="perPage" v-model="currentPage" prev-text="이전" next-text="다음" hide-goto-end-buttons align="center"/>
    </nav>
  </b-card>
</template>

<script>

export default {
  inheritAttrs: false,
  props: {
    caption: {
      type: String,
      default: 'Table'
    },
    hover: {
      type: Boolean,
      default: false
    },
    striped: {
      type: Boolean,
      default: true
    },
    bordered: {
      type: Boolean,
      default: true
    },
    small: {
      type: Boolean,
      default: false
    },
    fixed: {
      type: Boolean,
      default: false
    },
    tableData: {
      type: [Array, Function],
      default: () => []
    },
    fields: {
      type: [Array, Object],
      default: () => []
    },
    perPage: {
      type: Number,
      default: 15,
    },
    dark: {
      type: Boolean,
      default: false
    }
  },
  data: () => {
    return {
      currentPage: 1,    
      Id:-1,
      confirm:'',
      complete:''  
    }
  },
  computed: {
    items: function() {
      const items =  this.tableData
      return Array.isArray(items) ? items : items()
    },
    totalRows: function () { return this.getRowCount() },
    captions: function() { return this.fields }
  },
  methods: {
    getBadge (status) {
      return status === 'Active' ? 'success'
        : status === 'Inactive' ? 'secondary'
          : status === 'Pending' ? 'warning'
            : status === 'Banned' ? 'danger' : 'primary'
    },
    getRowCount: function () {
      return this.items.length
    },
    onClick(Arg_Id, Arg_confirm , Arg_complete){      
      this.Id = Arg_Id
      this.confirm = Arg_confirm
      this.complete = Arg_complete

      const URI = 'http://172.30.1.46:8091/depositUpdate'
      this.$axios.post(`${URI}`,
        {
          "Id" : this.Id,
          "confirm" : this.confirm,
          "complete" : this.complete
        }
      )
      .then((result) => {         
          const baseURI = 'http://172.30.1.46:8091/depositCheckList'
          this.$axios.post(`${baseURI}`)
          .then((result) => {                        
            if(result.data != null) {
              this.tableData = result.data      
            }else if(result.data == null){
              this.tableData = []
            }
          })
      })
    },
    onCompleted(Arg_Id, Arg_confirm, Arg_complete, Arg_WalletAddress, Arg_BaseUnit, Arg_BaseAmount, Arg_TargetUnit,
      Arg_TargetAmount, Arg_RateByBase , Arg_Transactions){
          
          //지갑 유효성 확인
          this.$axios.post('http://testnet.dubu4.com:6552',
            {
              "jsonrpc":"2.0",
              "id":1,
              "method":"validateaddress",
              "params":[Arg_WalletAddress]
            })
            .then((result) => {                     
              if(result.data.result === true){

                //데이터 전송
                this.$axios.post('http://172.30.1.46:8091/complete',
                {
                  "BuySell" : "Buy",
                  "Id" : Arg_Id ,
                  "WalletAddress" : Arg_WalletAddress,
                  "BaseUnit" : Arg_BaseUnit,
                  "BaseAmount" : Arg_BaseAmount,
                  "TargetUnit" : Arg_TargetUnit,
                  "TargetAmount" : Arg_TargetAmount,
                  "RateByBase" : Arg_RateByBase,
                  "Transactions" : Arg_Transactions
                })        
                .then((result) => {
                  
                })

                //송금처리
                this.$axios.post('http://testnet.dubu4.com:6552',
                {
                  "jsonrpc":"2.0",
                  "id":1,
                  "method":"sendtoaddress",
                  "params":[Arg_WalletAddress, Arg_TargetAmount*1000000]
                })
                .then((result) => {
                  
                })

                // 승인여부 변경 및 페이지 전환
                this.Id = Arg_Id
                this.confirm = Arg_confirm
                this.complete = Arg_complete

                const URI = 'http://172.30.1.46:8091/depositUpdate'
                this.$axios.post(`${URI}`,
                  {
                    "Id" : this.Id,
                    "confirm" : this.confirm,
                    "complete" : this.complete
                  }
                )
                .then((result) => {         
                    const baseURI = 'http://172.30.1.46:8091/depositCheckList'
                    this.$axios.post(`${baseURI}`)
                    .then((result) => {                        
                      if(result.data != null) {
                        this.tableData = result.data      
                      }else if(result.data == null){
                        this.tableData = []
                      }
                    })
                })

              }else {
                alert('사용자 지갑 주소가 유효하지 않습니다.')
              }
            })      

    }
  }
}
</script>
