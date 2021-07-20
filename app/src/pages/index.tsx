import {
  chakra,
  Button,
  Stack,
} from "@chakra-ui/react"
import Link from './components/Link';
import React from "react"
import MyPagesTemplate from "./myPagesTemplate";
const Home = (): JSX.Element => {
  return (
    <>
      <chakra.div><MyPagesTemplate /> ユーザー登録して一緒に記事を投稿しましょう！
        <Stack direction="row" align="center">
          <Link href="/registration">
            <Button colorScheme="teal" variant="solid">
              ユーザー登録
            </Button>
          </Link>
        </Stack>
      </chakra.div>
    </>
  )
}

export default Home