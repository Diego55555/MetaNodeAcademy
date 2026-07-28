// components/Card.tsx
import React from 'react'

type CardProps = {
  children: React.ReactNode // ReactNode 类型允许任何可渲染的内容
}

export default function Card({ children }: CardProps) {
  return (
    <div style={{ border: '1px solid #ccc', padding: '16px', borderRadius: '8px' }}>
      <div className="font-bold text-lg mb-2">
        卡片主标题
      </div>
      {children}
    </div>
  )
}