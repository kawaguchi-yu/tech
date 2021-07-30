import {
	Menu,
	MenuButton,
	MenuList,
	MenuItem,
	Box,
	Flex,
	Button,
	useColorModeValue,
	Spacer,
	Heading,
	Image,
} from '@chakra-ui/react';
import Link from './components/Link';
import React, { useState, useEffect } from "react"
type dataStruct = {
	ID: string
	CreatedAt: string
	UpdatedAt: string
	DeletedAt: string
	Name: string
	EMail: string
	Password: string
	Posts: string
	Profile: string
	ProfileID: string
	Goods: string
};
var returnData: dataStruct = {
	ID: "",
	CreatedAt: "",
	UpdatedAt: "",
	DeletedAt: "",
	Name: "",
	EMail: "",
	Password: "",
	Posts: "",
	Profile: "",
	ProfileID: "",
	Goods: "",
}
const Template = () => {

	useEffect(() => {
		fetch("http://localhost:8080/user", {
			mode: "cors",
			method: "GET",
			headers: { "Content-Type": "application/json", }, // JSON形式のデータのヘッダー
			credentials: 'include',
		}).then((res) => res.json())
			.then((data) => {
				const result = JSON.stringify(data)
				const result2: dataStruct = JSON.parse(result)
				returnData = result2
				setUser(returnData)
				if (returnData == null) {
					console.log("データはないよ！", returnData)
				} else {
					setHasCookie(true)
					console.log("データはあるよ！", returnData)
				}
			})
		fetch("http://localhost:8080/icon", {
			mode: "cors",
			method: "GET",
			credentials: 'include',
		}).then((res) => res.blob())
			.then((data) => {
				setIcon(data);
				console.log("データ", data)
			})
	}, [])
	const [user, setUser] = useState<dataStruct>(returnData);
	const [icon, setIcon] = useState<Blob>()
	const [hasCookie, setHasCookie] = useState<boolean>(false);
	return (
		<>
			<Flex bg={useColorModeValue('gray.100', 'gray.900')} alignItems={'center'}>
				<Link href="/">
					<Box h={16} p={2} color="Highlight">
						<Heading>Techer</Heading>
					</Box>
				</Link>
				<Spacer />
				<Box mr={4}>
					<Flex direction="row" align="center">
						{hasCookie
							? <><Heading mr="4">welcome!{user.Name}</Heading>
								<Spacer />
								<Menu>
									<MenuButton as={Button} h={16} p={2}>
										{icon && <Image boxSize="50px"
											borderRadius="full"
											src={(window.URL.createObjectURL(icon))}
											alt="select picture" />}
									</MenuButton>
									<MenuList>
										<Link href="/myPages"><MenuItem>マイページ</MenuItem></Link>
										<Link href="/post"><MenuItem>クイズを投稿する</MenuItem></Link>
										<Link href="/config"><MenuItem>設定</MenuItem></Link>
										<Link href="/terms"><MenuItem>利用規約</MenuItem></Link>
										<MenuItem>ログアウト</MenuItem>
									</MenuList>
								</Menu>
							</>
							: <>
								<Link href="/signup">
									<Button mr="4" colorScheme="teal" variant="solid">
										ユーザー登録
									</Button>
								</Link>
								<Link href="/login">
									<Button colorScheme="teal">
										ログイン
									</Button>
								</Link>
							</>
						}
					</Flex>
				</Box>
			</Flex>
		</>
	);
}
export default Template