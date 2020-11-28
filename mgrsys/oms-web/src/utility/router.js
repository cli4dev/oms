import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'menu',
      component: () => import('../pages/menu/menu.vue'),
      children:[
        
				{
					path: 'oms/down/shelf',
					name: 'omsDownShelf',
					component: () => import('../pages/oms/down/shelf.vue')
				},
				{
					path: 'oms/down/shelf.view',
					name: 'omsDownShelfView',
					component: () => import('../pages/oms/down/shelf.view.vue')
				},
				{
					path: 'oms/product/line',
					name: 'omsProductLine',
					component: () => import('../pages/oms/product/line.vue')
				},
				{
					path: 'oms/product/line.view',
					name: 'omsProductLineView',
					component: () => import('../pages/oms/product/line.view.vue')
				},
				{
					path: 'vds/order/query',
					name: 'vdsOrderQuery',
					component: () => import('../pages/vds/order/query.vue')
				},
				{
					path: 'vds/order/query.view',
					name: 'vdsOrderQueryView',
					component: () => import('../pages/vds/order/query.view.vue')
				},
				{
					path: 'vds/channel/error/code',
					name: 'vdsChannelErrorCode',
					component: () => import('../pages/vds/channel/error_code.vue')
				},
				{
					path: 'vds/channel/error/code.view',
					name: 'vdsChannelErrorCodeView',
					component: () => import('../pages/vds/channel/error_code.view.vue')
				},
				{
					path: 'dds/dictionary/info',
					name: 'ddsDictionaryInfo',
					component: () => import('../pages/dds/dictionary/info.vue')
				},
				{
					path: 'vds/order/exp',
					name: 'vdsOrderExp',
					component: () => import('../pages/vds/order/exp.vue')
				},
				{
					path: 'vds/order/exp.view',
					name: 'vdsOrderExpView',
					component: () => import('../pages/vds/order/exp.view.vue')
				},
				{
					path: 'oms/order/delivery',
					name: 'omsOrderDelivery',
					component: () => import('../pages/oms/order/delivery.vue')
				},
				{
					path: 'oms/order/delivery.view',
					name: 'omsOrderDeliveryView',
					component: () => import('../pages/oms/order/delivery.view.vue')
				},
				{
					path: 'beanpay/upaccount/info',
					name: 'beanpayUpAccountInfo',
					component: () => import('../pages/beanpay/account/up/info.vue')
				},
				{
					path: 'beanpay/upaccount/info.view',
					name: 'beanpayUpAccountInfoView',
					component: () => import('../pages/beanpay/account/up/info.view.vue')
				},
				{
					path: 'beanpay/downaccount/info',
					name: 'beanpayDownAccountInfo',
					component: () => import('../pages/beanpay/account/down/info.vue')
				},
				{
					path: 'beanpay/downaccount/info.view',
					name: 'beanpayDownAccountInfoView',
					component: () => import('../pages/beanpay/account/down/info.view.vue')
				},
				{
					path: 'oms/up/shelf',
					name: 'omsUpShelf',
					component: () => import('../pages/oms/up/shelf.vue')
				},
				{
					path: 'oms/up/shelf.view',
					name: 'omsUpShelfView',
					component: () => import('../pages/oms/up/shelf.view.vue')
				},
				{
					path: 'oms/down/product',
					name: 'omsDownProduct',
					component: () => import('../pages/oms/down/product.vue')
				},
				{
					path: 'oms/down/product.view/:no',
					name: 'omsDownProductView',
					component: () => import('../pages/oms/down/product.view.vue')
				},
				{
					path: 'vds/order/info',
					name: 'vdsOrderInfo',
					component: () => import('../pages/vds/order/info.vue')
				},
				{
					path: 'vds/order/info.view',
					name: 'vdsOrderInfoView',
					component: () => import('../pages/vds/order/info.view.vue')
				},
				{
					path: 'beanpay/upaccount/record',
					name: 'beanpayUpAccountRecord',
					component: () => import('../pages/beanpay/account/up/record.vue')
				},
				{
					path: 'beanpay/upaccount/record.view',
					name: 'beanpayUpAccountRecordView',
					component: () => import('../pages/beanpay/account/up/record.view.vue')
				},
				{
					path: 'beanpay/downaccount/record',
					name: 'beanpayDownAccountRecord',
					component: () => import('../pages/beanpay/account/down/record.vue')
				},
				{
					path: 'beanpay/downaccount/record.view',
					name: 'beanpayDownAccountRecordView',
					component: () => import('../pages/beanpay/account/down/record.view.vue')
				},
				{
					path: 'oms/refund/up/return',
					name: 'omsRefundUpReturn',
					component: () => import('../pages/oms/refund/up_return.vue')
				},
				{
					path: 'oms/refund/up/return.view',
					name: 'omsRefundUpReturnView',
					component: () => import('../pages/oms/refund/up_return.view.vue')
				},
				{
					path: 'vds/order/notify',
					name: 'vdsOrderNotify',
					component: () => import('../pages/vds/order/notify.vue')
				},
				{
					path: 'vds/order/notify.view',
					name: 'vdsOrderNotifyView',
					component: () => import('../pages/vds/order/notify.view.vue')
				},
				{
					path: 'oms/refund/info',
					name: 'omsRefundInfo',
					component: () => import('../pages/oms/refund/info.vue')
				},
				{
					path: 'oms/refund/info.view',
					name: 'omsRefundInfoView',
					component: () => import('../pages/oms/refund/info.view.vue')
				},
				{
					path: 'oms/up/product',
					name: 'omsUpProduct',
					component: () => import('../pages/oms/up/product.vue')
				},
				{
					path: 'oms/up/product.view/:no',
					name: 'omsUpProductView',
					component: () => import('../pages/oms/up/product.view.vue')
				},
				{
					path: 'oms/up/channel',
					name: 'omsUpChannel',
					component: () => import('../pages/oms/up/channel.vue')
				},
				{
					path: 'oms/up/channel.view',
					name: 'omsUpChannelView',
					component: () => import('../pages/oms/up/channel.view.vue')
				},
				{
					path: 'lcs/life/time',
					name: 'lcsLifeTime',
					component: () => import('../pages/lcs/life/time.vue')
				},
				{
					path: 'lcs/life/time.view',
					name: 'lcsLifeTimeView',
					component: () => import('../pages/lcs/life/time.view.vue')
				},
				{
					path: 'oms/notify/info',
					name: 'omsNotifyInfo',
					component: () => import('../pages/oms/notify/info.vue')
				},
				{
					path: 'oms/notify/info.view',
					name: 'omsNotifyInfoView',
					component: () => import('../pages/oms/notify/info.view.vue')
				},
				{
					path: 'oms/audit/info',
					name: 'omsAuditInfo',
					component: () => import('../pages/oms/audit/info.vue')
				},
				{
					path: 'oms/audit/info.view',
					name: 'omsAuditInfoView',
					component: () => import('../pages/oms/audit/info.view.vue')
				},
				{
					path: 'vds/channel/info',
					name: 'vdsChannelInfo',
					component: () => import('../pages/vds/channel/info.vue')
				},
				{
					path: 'vds/channel/info.view',
					name: 'vdsChannelInfoView',
					component: () => import('../pages/vds/channel/info.view.vue')
				},
				{
					path: 'oms/down/channel',
					name: 'omsDownChannel',
					component: () => import('../pages/oms/down/channel.vue')
				},
				{
					path: 'oms/down/channel.view',
					name: 'omsDownChannelView',
					component: () => import('../pages/oms/down/channel.view.vue')
				},
				{
					path: 'oms/order/info',
					name: 'omsOrderInfo',
					component: () => import('../pages/oms/order/info.vue')
				},
				{
					path: 'oms/order/info.view',
					name: 'omsOrderInfoView',
					component: () => import('../pages/oms/order/info.view.vue')
				},
				{
					path: 'tsk/system/task',
					name: 'tskSystemTask',
					component: () => import('../pages/tsk/system/task.vue')
				},
				{
					path: 'tsk/system/task.view',
					name: 'tskSystemTaskView',
					component: () => import('../pages/tsk/system/task.view.vue')
				},
				{
					path: 'oms/canton/info',
					name: 'omsCantonInfo',
					component: () => import('../pages/oms/canton/info.vue')
				},
				{
					path: 'report/down/channel',
					name: 'reportDownChannel',
					component: () => import('../pages/report/downchannel.vue')
				},
				{
					path: 'report/up/channel',
					name: 'reportUpChannel',
					component: () => import('../pages/report/upchannel.vue')
				},
				{
					path: 'report/profit',
					name: 'reportProfit',
					component: () => import('../pages/report/profit.vue')
				},
      ]
	},
	{
      path:'/ssocallback',
      name:'ssocallback',
      component: () => import('../pages/login/ssocallback.vue'),
    }
  ]
})