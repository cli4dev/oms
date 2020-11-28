
export default {
    toDetail(orderNo) {
      this.$router.push({name:'OrderDetail', params:{order_no: orderNo}})
    },
  
    dealwithNum(result){
      result = result.trim()
      if( result.split('.').length > 1 ){
        let integer = result.split('.')[0]
        let decimal = result.split('.')[1]
        let len = integer.length
        if(len <= 3){
          if(integer === ""){
            return "0." + decimal
          } else {
            return result
          }
        } else {
          let r = len % 3
          result = r > 0? integer.slice(0,r) + "," + integer.slice(r,len).match(/\d{3}/g).join(",") : integer.slice(r,len).match(/\d{3}/g).join(",")
          return result + '.' + decimal
        }
      }
      if( result.split('.').length === 1){
        let len = result.length
        if(len <= 3){
          return result
        } else {
          let r = len % 3
          return r > 0? result.slice(0,r) + "," + result.slice(r,len).match(/\d{3}/g).join(",") : result.slice(r,len).match(/\d{3}/g).join(",")
        }
      }
      if( result.length === 0){
        return '-'
      }
    }
  }