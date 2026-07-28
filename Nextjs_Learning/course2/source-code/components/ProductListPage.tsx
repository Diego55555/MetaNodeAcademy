// 默认是服务端组件，因为数据是静态的，不需要在每次请求时都重新渲染。
// 但是因为它的父组件 HomePage 是一个客户端组件，所以它在这里也变成了客户端组件。
const products = [
  { id: 1, name: '笔记本电脑' },
  { id: 2, name: '智能手机' },
  { id: 3, name: '无线耳机' },
]

export default function ProductListPage() {
  return (
    <div className="bg-gray-100 p-4 rounded-md">
      <h2>商品列表</h2>
      <ul>
        {products.map((product, idx) => (
          <li key={product.id}>
            {product.name}
          </li>
        ))}
      </ul>
    </div>
  )
}