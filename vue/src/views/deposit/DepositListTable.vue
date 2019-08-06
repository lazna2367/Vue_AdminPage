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
      <template slot="Confirm" slot-scope="data">  
        <!-- pending -->
        <b-dropdown v-if="data.item.Confirm == 'not'" size="sm" id="ddown_secondary" variant="secondary" class="m-1" text="미확인">        
          <b-dropdown-item @click="onClick(data.item.Id,'not','not')">미확인</b-dropdown-item>
          <b-dropdown-item @click="onClick(data.item.Id,'pending','not')">보류</b-dropdown-item>
          <b-dropdown-item @click="onClick(data.item.Id,'confirmed','not')">확인</b-dropdown-item>
        </b-dropdown>        
        <b-dropdown v-if="data.item.Confirm == 'pending'" size="sm" id="ddown_warning" variant="warning" class="m-1" text="대기중">        
          <b-dropdown-item @click="onClick(data.item.Id,'not','not')">미확인</b-dropdown-item>
          <b-dropdown-item @click="onClick(data.item.Id,'pending','not')">보류</b-dropdown-item>
          <b-dropdown-item @click="onClick(data.item.Id,'confirmed','not')">확인</b-dropdown-item>
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
      //console.log('돈다돌아' + Id + ' / ' +  confirm + ' / ' + complete);      
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
        console.log(result)               
          const baseURI = 'http://172.30.1.46:8091/depositList'
          this.$axios.post(`${baseURI}`)
          .then((result) => {            
            console.log(result)    
            if(result.data != null){
              this.tableData = result.data      
            } else if (result.data == null){
              this.tableData = []
            }
          })

        //this.axiosload()
      })
    }
  }
}
</script>
