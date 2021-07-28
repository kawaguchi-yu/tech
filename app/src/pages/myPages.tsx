import React, { useState, useEffect } from "react"
import {
	chakra,
	Stack,
	Button,
	Heading,
	Box,
	VStack,
} from '@chakra-ui/react';
import Link from './components/Link';
import Template from "./template";
type user = {
	ID: number
	CreatedAt: string
	UpdatedAt: string
	DeletedAt: string
	UserID: number
	Title: string
	WrongAnswer1: string
	WrongAnswer2: string
	WrongAnswer3: string
	Explanation: string
	Tags: any
	Goods: any
};

const MyPages = (): JSX.Element => {
	useEffect(() => {
		fetch("http://localhost:8080/getuserpost", {
			mode: "cors",
			method: "POST",
			headers: { "Content-Type": "application/json", }, // JSON形式のデータのヘッダー
			credentials: 'include',
		})
			.then((res) => res.json())
			.then((data) => {
				data.forEach(array =>
					setPostDatas(postDatas => [...postDatas, array]),
					console.log(data)
				)
			})
	}, [])
	const [postDatas, setPostDatas] = useState<user[]>([])
	const PostsView = () => {
		if(postDatas[0]){
		var dayData = postDatas[0].CreatedAt
		var dayData2 = dayData.substring(0,10)
		}
		return (<>
			<VStack>
				{postDatas[0] && <><Box>{dayData2}</Box>
					<Box><Link href="">{postDatas[0].Title}</Link></Box>
				</>}
			</VStack>
		</>)
	}
	return (
		<>
			<Template />
			<chakra.div>
				<Stack direction="row" align="center">
					<PostsView />
				</Stack>
			</chakra.div>
		</>
	)
}
export default MyPages