import {
	Menu,
	MenuButton,
	MenuList,
	MenuItem,
	Box,
	Grid,
	Flex,
	Button,
	useColorModeValue,
	Spacer,
	Heading,
	Image,
	Input,
	InputGroup,
	InputRightElement,
	IconButton,
	FormControl,
	Stack,
	HStack,
	VStack,
} from '@chakra-ui/react';
import { Search2Icon, HamburgerIcon } from '@chakra-ui/icons'
import Link from "./components/Link"
import { useRouter } from 'next/router'
import React, { useState, useEffect } from "react"
import { useForm } from "react-hook-form";
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
type Form = {
	Word: string
}
const SearchFunc = () => {
	const { register, formState, formState: { errors }, getValues } = useForm<Form>({
		mode: "onChange"
	});
	const router = useRouter()
	const Search = () => {
		const word = getValues("Word")
		if (word == "") {
			return
		}
		router.push({
			pathname: `/search/`,
			query: { word }
		})
	}
	return (<>
		<FormControl isInvalid={errors.Word ? true : false}>
			<InputGroup>
				<Input
					{...register("Word", {
						required: true,
						pattern: {
							value: /^[^^＾"”`‘'’<>＜＞_＿%$#＆％＄|￥]{1,20}$/,
							message: '特殊文字を使用しないでください' // JS only: <p>error message</p> TS only support string
						}
					})}
				/>
				<InputRightElement>
					<IconButton
						aria-label="Search database"
						icon={<Search2Icon />}
						onClick={Search}
						disabled={!formState.isValid} />
				</InputRightElement>
			</InputGroup>
			{errors.Word && errors.Word.message}
		</FormControl>
	</>)
}
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
	const GuestLogin = () => {
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
		<Flex bg={useColorModeValue("blue.100", 'gray.900')} alignItems={'center'}>
			<Link href="/">
				<Box height={16} p={2} color="honeydew">
					<Heading>Techer</Heading>
				</Box>
			</Link>
			<Spacer />
			<Menu>
				<MenuButton>
					<IconButton
						margin="2"
						aria-label="Search database"
						icon={<Search2Icon />}
						colorScheme="white" />
				</MenuButton>
				<MenuList>
					<SearchFunc />
				</MenuList>
			</Menu>
			<Flex>
				{user
					? <Menu>
						<MenuButton as={Button} height={16} width={16} p={2} rounded="full">
							{user.IconBlob && <Image
								boxSize="50px"
								src={(window.URL.createObjectURL(user.IconBlob))}
								alt="userIcon"
								fallbackSrc="https://via.placeholder.com/150" />}
						</MenuButton>
						<MenuList>
							<Link href={`/${user.Name}`}><MenuItem>マイページ</MenuItem></Link>
							<Link href="/post"><MenuItem>クイズを投稿する</MenuItem></Link>
							<Link href="/config"><MenuItem>設定</MenuItem></Link>
							<Link href="/terms"><MenuItem>利用規約</MenuItem></Link>
							<MenuItem onClick={Logout}>ログアウト</MenuItem>
						</MenuList>
					</Menu>
					: <>
						<Stack display={{ base: "block", md: "none" }} margin="2">
							<Menu>
								<MenuButton>
									<HamburgerIcon />
								</MenuButton>
								<MenuList>
									<MenuItem onClick={GuestLogin}>ゲストログイン</MenuItem>
									<Link href="/signup"><MenuItem>ユーザー登録</MenuItem></Link>
									<Link href="/login"><MenuItem>ログイン</MenuItem></Link>
								</MenuList>
							</Menu>
						</Stack>
						<Button
							margin="2"
							display={{ base: "none", md: "block" }}
							colorScheme="linkedin"
							onClick={GuestLogin}>
							ゲストログイン
						</Button>
						<Link href="/signup">
							<Button
								margin="2"
								display={{ base: "none", md: "block" }}
								colorScheme="teal">
								ユーザー登録
							</Button>
						</Link>
						<Link href="/login">
							<Button
								margin="2"
								display={{ base: "none", md: "block" }}
								colorScheme="teal">
								ログイン
							</Button>
						</Link>
					</>
				}
			</Flex>
		</Flex>
	);
}
export default Template