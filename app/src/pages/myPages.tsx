import React, { useState, useEffect } from "react"
import {
    chakra,
    Stack,
    Button,
    Heading,
} from '@chakra-ui/react';
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
    EMail: "a",
    Password: "",
    Posts: "",
    Profile: "",
    ProfileID: "",
    Goods: "",
}

const MyPages = (): JSX.Element => {

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


                // userData.ID = result[0] 
                setEMail(userData)
            })
    }, [])
    const [email, setEMail] = useState<user>({ ID: "", CreatedAt: "", UpdatedAt: "", DeletedAt: "", Name: "", EMail: "", Password: "", Posts: "", Profile: "", ProfileID: "", Goods: "", });
    return (
        <>
            <chakra.div>
                <Stack direction="row" align="center">
                    <Button colorScheme="teal" variant="solid">
                        ユーザー登録
                    </Button>
                    <Button type="submit"
                        colorScheme="teal">
                        ログイン
                    </Button>
                    <Heading>welcome!{email.Name}</Heading>
                </Stack>
            </chakra.div>
        </>
    )
}
export default MyPages