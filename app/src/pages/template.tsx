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
import Link from "./components/Link"
import { useRouter } from 'next/router'
import React, { useState, useEffect } from "react"
type user = {
	ID: string
	CreatedAt: string
	UpdatedAt: string
	DeletedAt: string
	Name: string
	EMail: string
	Password: string
	Profile: string
	ProfileID: string
	Goods: string
	Icon: string
	IconBlob: Blob
};
const Template = () => {
	const router = useRouter()
	const [user, setUser] = useState<user>();
	useEffect(() => {
		fetch("http://localhost:8080/user", {
			mode: "cors",
			method: "GET",
			credentials: 'include',
		}).then((res) => res.json())
			.then((data) => {
				const userData: user = data
				if (userData == null) {
					console.log("データはないよ！", userData)
				} else {
					let bin = atob(userData.Icon.replace(/^.*,/, ''));
					let buffer = new Uint8Array(bin.length);
					for (let i = 0; i < bin.length; i++) {
						buffer[i] = bin.charCodeAt(i);
					} let blob = new Blob([buffer.buffer], {
						type: "image/jpeg"
					});
					userData.IconBlob = blob
					setUser(userData)
				}
			}).catch(() => {
				console.error("データを貰ってくることができませんでした")
			})
	}, [])
	const GuestLogin = () =>{
		fetch("http://localhost:8080/guestlogin", {
			mode: "cors",
			method: "GET",
			credentials: 'include',
		}).then((res) => res.json())
			.then(() => {
				router.reload()
			})
	}
	const Logout = () => {
		fetch("http://localhost:8080/logout", {
			mode: "cors",
			method: "GET",
			credentials: 'include',
		}).then((res) => res.json())
			.then(() => {
				router.reload()
			})
	}
	return (
		<>
			<Flex bg={useColorModeValue("blue.100", 'gray.900')} alignItems={'center'}>
				<Link href="/">
					<Box h={16} p={2} color="Highlight">
						<Heading>Techer</Heading>
					</Box>
				</Link>
				<Spacer />
				<Box mr={4}>
					<Flex direction="row" align="center">
						{user
							? <><Heading mr="4">welcome!{user.Name}</Heading>
								<Spacer />
								<Menu>
									<MenuButton as={Button} h={16} p={2}>
										{user.IconBlob && <Image boxSize="50px"
											borderRadius="full"
											src={(window.URL.createObjectURL(user.IconBlob))}
											alt="select picture" />}
									</MenuButton>
									<MenuList>
										<Link href={`/${user.Name}`}><MenuItem>マイページ</MenuItem></Link>
										<Link href="/post"><MenuItem>クイズを投稿する</MenuItem></Link>
										<Link href="/config"><MenuItem>設定</MenuItem></Link>
										<Link href="/terms"><MenuItem>利用規約</MenuItem></Link>
										<MenuItem onClick={Logout}>ログアウト</MenuItem>
									</MenuList>
								</Menu>
							</>
							: <>
							<Button　mr="4" colorScheme="linkedin" onClick={GuestLogin}>
								ゲストログイン
							</Button>
								<Link href="/signup">
									<Button mr="4" colorScheme="teal">
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