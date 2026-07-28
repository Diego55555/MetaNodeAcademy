'use client'

import React, { useState } from 'react'
import Greeting from './Greeting'
import ConditionalRenderPage from './ConditionalRenderPage'
import ProductListPage from './ProductListPage'
import EffectDemoPage from './EffectDemoPage'
import CompositionPage from './CompositionPage'

// 这是一个客户端组件，因为它使用了hooks useState，而 useState只能在客户端组件中使用。而且它的所有子组件也必须是客户端组件。
export default function HomePage() {
    const [count, setCount] = useState(0)

    return (
        <div>
            <p>你点击了 {count} 次</p>
            <button
                className="bg-red-500 text-white p-2 rounded-md"
                onClick={() => setCount(count + 1)}
            >
                点我
            </button>
            <Greeting name="react" />
            <Greeting name="nextjs" />
            <ConditionalRenderPage />
            <ProductListPage />
            <EffectDemoPage/>
            <CompositionPage />
        </div>
    )
}
