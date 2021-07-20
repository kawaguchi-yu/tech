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
type user = {
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
var userData: user = {
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

const MyPagesTemplate = () => {
	useEffect(() => {
		fetch("http://localhost:8080/user", {
			mode: "cors",
			method: "GET",
			headers: { "Content-Type": "application/json", }, // JSON形式のデータのヘッダー
			credentials: 'include',
		})
			.then((res) => res.json())
			.then((data) => {
				const result = JSON.stringify(data)
				const result2: user = JSON.parse(result)
				userData = result2
				setEMail(userData)
				if (userData == null) {
					console.log("データはないよ！", userData)
				} else {
					setHasCookie(true)
					console.log("データはあるよ！", userData)
				}
			})
	}, [])
	const [email, setEMail] = useState<user>({ ID: "", CreatedAt: "", UpdatedAt: "", DeletedAt: "", Name: "", EMail: "", Password: "", Posts: "", Profile: "", ProfileID: "", Goods: "", });
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
							? <><Heading mr="4">welcome!{email.Name}</Heading>
								<Spacer />
								<Menu>
									<MenuButton as={Button} h={16} p={2}>
										<Image
											boxSize="50px"
											borderRadius="full"
											src="https://bit.ly/sage-adebayo"
											alt="Segun Adebayo" />
									</MenuButton>
									<MenuList>
										<Link href="/config"><MenuItem>設定</MenuItem></Link>
										<Link href="/terms"><MenuItem>利用規約</MenuItem></Link>
										<MenuItem>ログアウト</MenuItem>
									</MenuList>
								</Menu>
							</>
							: <>
								<Link href="/registration">
									<Button mr="4" colorScheme="teal" variant="solid">
										ユーザー登録
									</Button>
								</Link>
								<Link href="/login">
									<Button colorScheme="teal">
										アイコン
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
export default MyPagesTemplate