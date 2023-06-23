import React from 'react'

type ContainerProps = {
  children: React.ReactNode
}
const Container = (props: ContainerProps) => {
  return <div className="container mx-auto px-4 md:px-32">{props.children}</div>
}

export default Container
