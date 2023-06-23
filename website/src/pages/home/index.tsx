import Container from '@/shared/layout/container'
import Button from '@/shared/ui/button'
import Editor from '@/shared/ui/editor'
import Header from '@/shared/ui/header'

import { useState } from 'react'

const Search = () => {
  const [searchHash, setSearchHash] = useState('')

  return (
    <input
      type="text"
      name="search"
      className="inline-flex w-full items-start rounded-2xl bg-grey px-6 py-5 font-fira-code text-base font-medium focus:outline-none md:w-auto md:px-10 md:text-lg"
      placeholder="Search paste with hash..."
      onChange={(e) => setSearchHash(e.target.value)}
    />
  )
}

function HomePage() {
  const [body, setBody] = useState('')

  return (
    <main>
      <Container>
        <Header />
        <Search />
        <Editor value={body} onChange={setBody} />
        <div className="flex justify-end">
          <div className="w-full md:w-1/6">
            <Button>Create</Button>
          </div>
        </div>
      </Container>
    </main>
  )
}

export default HomePage
