
import Vue from "vue"
import Common from './common'
import DateConvert from './date'

Vue.filter('DateFilter', (value, format) => {
  let res
  if (value === '') {
    return '-'
  } else {
    res = DateConvert(format, value)
    return res
  }
})

Vue.filter('StringFilter', value => {
  if (value === '') {
    return '---'
  }else{
    return value
  }
})

Vue.filter('EllipsisFilter', (value, number) => {
  if (value) {
    if (value.length <= number) {
      return value
    }
    else {
      let subval = value.slice(0, number - 1) + '...'
      return subval
    }
  }
  else {
    return '-'
  }
})

Vue.filter('AmountFilter', (value)=>{
  if (value == null||value == undefined ||value == "") {
    return 0
  }
  value = Math.round(value*100)/100

  var c = (value.toString().indexOf ('.') !== -1) ? value.toLocaleString() : value.toString().replace(/(\d)(?=(?:\d{3})+$)/g, '$1,');
  return c
})

Vue.filter('FeeFilter', (value)=>{
  if (value == null||value == undefined ||value == "") {
    return 0
  }
  value = Math.round(value*1000)/1000

  return value
})
