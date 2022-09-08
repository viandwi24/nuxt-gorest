import { Brand } from '~/types'

export const brands: Brand[] = [
  {
    id: 1,
    name: 'McDonalds',
    image:
      'https://res.cloudinary.com/wnotw/images/c_limit,w_1536,q_auto:good,f_auto/v1622443418/l6vzyybadqseihd7b6p7/mcdonald-s',
    rating: 4.7,
    avearageTime: 20,
    menus: [
      {
        id: 1,
        restaurantId: 1,
        name: 'Big Mac',
        category: 'Burgers',
        price: 5.99,
        description:
          'Big Mac is a hamburger sold by international fast food restaurant chain McDonald’s.',
        image:
          'https://res.cloudinary.com/wnotw/images/c_limit,w_1536,q_auto:good,f_auto/v1622443418/l6vzyybadqseihd7b6p7/mcdonald-s',
      },
      {
        id: 2,
        restaurantId: 1,
        name: 'Quarter Pounder with Cheese',
        category: 'Burgers',
        price: 4.99,
        description:
          'The Quarter Pounder with Cheese is a hamburger sold by international fast food restaurant chain McDonald’s.',
        image:
          'https://res.cloudinary.com/wnotw/images/c_limit,w_1536,q_auto:good,f_auto/v1622443418/l6vzyybadqseihd7b6p7/mcdonald-s',
      },
    ],
  },
  {
    id: 2,
    name: 'KFC',
    image:
      'https://cdnwpedutorenews.gramedia.net/wp-content/uploads/2022/02/02090207/KFC_0.png',
    rating: 4.5,
    avearageTime: 20,
    menus: [
      {
        id: 3,
        restaurantId: 2,
        name: 'Original Recipe',
        category: 'Burgers',
        price: 4.99,
        description:
          'The Original Recipe is a fried chicken recipe developed by Harland Sanders, the founder of Kentucky Fried Chicken.',
        image:
          'https://cdnwpedutorenews.gramedia.net/wp-content/uploads/2022/02/02090207/KFC_0.png',
      },
    ],
  },
  {
    id: 3,
    name: 'Burger King',
    image:
      'https://ik.imagekit.io/tvlk/cul-asset/guys1L+Yyer9kzI3sp-pb0CG1j2bhflZGFUZOoIf1YOBAm37kEUOKR41ieUZm7ZJ/cul-assets-252301483284-b172d73b6c43cddb/culinary/asset/REST_bbq-960x720-FIT_AND_TRIM-37b8dff2d5a900882b150cb7d0efcddc.jpeg?tr=q-40,c-at_max,w-1080,h-1920&_src=imagekit',
    rating: 4.2,
    avearageTime: 20,
    menus: [],
  },
  {
    id: 4,
    name: 'Pizza Hut',
    image:
      'https://ik.imagekit.io/tvlk/cul-asset/guys1L+Yyer9kzI3sp-pb0CG1j2bhflZGFUZOoIf1YOBAm37kEUOKR41ieUZm7ZJ/cul-assets-252301483284-b172d73b6c43cddb/culinary/asset/REST_557-842x720-FIT_AND_TRIM-ebc88d9f46405c3e03e4208b0e18b88c.jpeg?tr=q-40,c-at_max,w-1080,h-1920&_src=imagekit',
    rating: 4.1,
    avearageTime: 10,
    menus: [],
  },
  {
    id: 5,
    name: 'Starbucks',
    image:
      'https://post.healthline.com/wp-content/uploads/2020/03/Starbucks_Storefront_732x549-thumbnail-732x549.jpg',
    rating: 4.4,
    avearageTime: 5,
    menus: [],
  },
  {
    id: 6,
    name: 'Dominos Pizza',
    image:
      'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSbT3OXYDyxKntinaaHuSg6yCERd2oxVEQYNQ&usqp=CAU',
    rating: 4.8,
    avearageTime: 15,
    menus: [],
  },
  {
    id: 7,
    name: 'Subway',
    image:
      'https://asset.kompas.com/crops/N84ZWhWErf6w53s911ga5LMwxE8=/0x92:1080x812/750x500/data/photo/2021/08/24/61246d40e2de7.jpg',
    rating: 4.2,
    avearageTime: 20,
    menus: [],
  },
  {
    id: 8,
    name: 'Bakmi GM',
    image:
      'https://res.cloudinary.com/wnotw/images/c_limit,w_1536,q_auto:good,f_auto/v1622443418/l6vzyybadqseihd7b6p7/mcdonald-s',
    rating: 4.5,
    avearageTime: 20,
    menus: [],
  },
]

export const getBrands = () => {
  return new Promise<Brand[]>((resolve) => {
    setTimeout(() => {
      resolve(brands)
    }, 2000)
  })
}

export const getBrand = (id: number) => {
  return new Promise<Brand | undefined>((resolve) => {
    setTimeout(() => {
      resolve(brands.find((brand) => brand.id === id))
    }, 0)
  })
}
