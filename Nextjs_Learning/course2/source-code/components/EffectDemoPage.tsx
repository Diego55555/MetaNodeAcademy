
import { useState, useEffect } from 'react'

export default function EffectDemoPage() {
  const [count, setCount] = useState(0)

  // 注意：这个组件是客户端组件，因为它使用了 useEffect 钩子和浏览器 api 也就是document.title
  // 客户端组件可以访问浏览器 API，如 document.title，服务端组件不能访问浏览器 API，因为它们在服务器上运行
  // 每次 count 变化时，更新浏览器标题
  useEffect(() => {
    document.title = `你点击了 ${count} 次`
  }, [count]) // 依赖数组，只有 count 变化时才执行

  return (
    <div>
      <p className="text-lg font-bold mb-4">查看浏览器标签页标题的变化</p>
      <button className="bg-blue-500 text-white p-2 rounded-md" onClick={() => setCount(count + 1)}>
        增加计数
      </button>
    </div>
  )
}