import { chakra } from "@chakra-ui/react"
import Link from 'next/link'
import React from "react"

const Home = ():JSX.Element => {
  return <>
        <chakra.div>ユーザー登録して一緒に記事を投稿しましょう！</chakra.div>
        <Link href="/test">
        <a>登録フォームに行く</a>
        </Link>
  </>
}

export default Home