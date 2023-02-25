import OrderCard from '../components/ordercard'
import Header from '../components/header'
import { Center, VStack, Text, Box, SimpleGrid } from '@chakra-ui/react';
import { useState } from 'react';
import { useNavigate } from "react-router-dom";
import { auth } from '../../firebase'
import { useAuthState } from 'react-firebase-hooks/auth';
import { onAuthStateChanged, signOut } from 'firebase/auth';

const dummyUser = {
  id: '129445',
  mobileNumber: '19789',
  addresses: [
    {
      userAddr1: "38 Nanyang Cres",
      userAddr2: "Saraca Hall",
      userAddr3: "22-11-22",
      city: "Singapore West",
      postalCode: "636866",
    },
    {
      userAddr1: "38 UTown",
      userAddr2: "Prince George Park",
      userAddr3: "24-1",
      city: "Singapore South",
      postalCode: "666666",
    },
  ],
}

const dummyData = [
  {
    id: '01273048120',
    name: 'Puri Virakarn',
    title: "parcel name",
    userAddrs: ["38 UTown", "Prince George Park", "24-1", "Singapore South", "666666"],
    instructions: ["Leave Parcel at the door"],
    leaveParcel: true,
    status: "On the way",
    mobileNumber: '123446',
    emailAddr: "adsfljas@opijpo.co.sg",
  },
  {
    id: '10283490870',
    name: 'Puri Virakarn',
    title: "parcel name",
    userAddrs: ["38 UTown", "Prince George Park", "24-1", "Singapore South", "666666"],
    instructions: ["Leave Parcel at the door"],
    leaveParcel: true,
    status: "Picked up",
    mobileNumber: '083290',
    emailAddr: "name@gmail.addr",
  },
  {
    id: '2134235432',
    name: 'Puri Virakarn',
    title: "parcel name",
    userAddrs: ["38 Nanyang Cres", "Saraca Hall", "22-11-22", "Singapore West", "636866"],
    instructions: ["Leave Parcel at the door"],
    leaveParcel: true,
    status: "Picked up",
    mobileNumber: '083290',
    emailAddr: "loc@facebook.ai",
  },
  {
    id: '18379489234',
    name: 'Puri Virakarn',
    title: "parcel name",
    userAddrs: ["38 Nanyang Cres", "Saraca Hall", "22-11-22", "Singapore West", "636866"],
    instructions: ["Leave Parcel at the door"],
    leaveParcel: true,
    status: "Order created",
    mobileNumber: '083290',
    emailAddr: "asdlf@asf.com",
  },
];

export default function OrderedList() {
  const [user] = useAuthState(auth);
  const navigate = useNavigate();
  const [data, setData] = useState(dummyData);
  const logOut = () => {
    alert('logging out...');
    signOut(auth).then(() => {
      // Sign-out successful.
      navigate("/login")
    }).catch((error) => {
      // An error happened.
      alert(error);
    });
  }
  return (
    <h1>
      <Header currentUser={user} logOut={logOut} />
      <SimpleGrid spacingY={4} mt={6} columns={{ 'sm': 1, 'md': 3,'xl': 4 }}>
        {data.map((eachOrder, index) => (
          <Box boxShadow={'xl'} m='auto'>
            <OrderCard alldata={data} data={eachOrder} index={index} userData={dummyUser} setData={setData}></OrderCard>
          </Box>
        )
        )}
      </SimpleGrid>
    </h1>
  )
}