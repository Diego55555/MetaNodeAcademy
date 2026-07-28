
import { useState } from 'react'

export default function ConditionalRenderPage() {
  const [isVisible, setIsVisible] = useState(true)
  const [isVis, setIsVis] = useState(1)

  return (
    <div>
      <button className='bg-blue-500 text-white p-2 rounded-md' onClick={() => setIsVisible(!isVisible)}>
        切换显示/隐藏
      </button>
      {/* 推荐：使用三元运算符，当条件为 false 时返回 null，确保不会意外渲染任何内容 */}
      {isVisible ? <p>这段文字现在是可见的。</p> : null}
      <div>
        <button className='bg-blue-500 text-white p-2 rounded-md' onClick={() => setIsVis(isVis ? 0 : 1)}>
          测试&&运算符
        </button>
        {/* 为了对比，这里展示 && 的用法。注意前面讲解中提到的潜在问题。 */}
        {isVis && <p>这段文字现在是可见的。</p>}
      </div>
    </div>
  )
}