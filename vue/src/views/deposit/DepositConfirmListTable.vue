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
      <template slot="Confirm">  
        <!-- pending -->
        <b-button block variant="success">처리완료</b-button>
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
    }
  }
}
</script>
