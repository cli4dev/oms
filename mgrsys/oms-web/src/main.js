import "jquery"
import "bootstrap"
import Vue from 'vue'
import App from './App.vue'
import router from './utility/router'
import store from './utility/store'
import DateConvert from './utility/date'
import DateFilter from './utility/filter';
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import {ssoHttpConfig} from 'qxnw-sso';
import {EnumUtility,EnumFilter} from 'qxnw-enum';


var config = process.env;
var ssocfg = ssoHttpConfig(config.VUE_APP_API_URL, "localStorage", config.VUE_APP_SSO_LOGIN_HOST, config.VUE_APP_IDENT);

//将sso和http都挂在vue对象中，方便使用
Vue.prototype.$sso = ssocfg.sso; 
Vue.prototype.$http = ssocfg.http;
Vue.prototype.EnumUtility = new EnumUtility() // 枚举字典
Vue.prototype.DateConvert = DateConvert //日期格式转换

Vue.use(ElementUI);
Vue.config.productionTip = false
console.log("当前环境：", process.env.NODE_ENV)


new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')