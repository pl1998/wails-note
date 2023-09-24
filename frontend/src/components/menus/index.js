import MenuContext from './../menus/RightMenu.vue'
import { h, render } from 'vue'

// 维护一个菜单实例
let curInstance = null
let seed = 1
export const contextMenu = () => {
    console.log("测试虚拟挂载")
    // 创建一个临时的div，用于挂载我们的菜单  
    const container = document.createElement('div')
    // 获取body标签，用于挂载整个菜单
    const appendTo = document.body 
    // 传给组件的props
    const props = {
    }
    // 渲染虚拟节点
    const vnode = h(
        MenuContext,
        props
    )
    // vnode为需要渲染的虚拟节点，container为渲染的容器
    console.log(container)
    render(vnode, container)
}
