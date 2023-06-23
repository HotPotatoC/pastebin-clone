import { SVGProps } from 'react'
const AvatarPlaceholder = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    width={62}
    height={62}
    fill="none"
    {...props}
  >
    <g clipPath="url(#a)">
      <mask
        id="b"
        width={62}
        height={62}
        x={0}
        y={0}
        maskUnits="userSpaceOnUse"
        style={{
          maskType: 'luminance',
        }}
      >
        <path
          fill="#fff"
          d="M62 31C62 13.88 48.12 0 31 0 13.88 0 0 13.88 0 31c0 17.12 13.88 31 31 31 17.12 0 31-13.88 31-31Z"
        />
      </mask>
      <g mask="url(#b)">
        <path fill="#0A0310" d="M62 0H0v62h62V0Z" />
        <path
          fill="#0043FF"
          d="M31.317 8.993C14.407 11.67 2.87 27.55 5.548 44.46 8.227 61.37 24.106 72.908 41.016 70.23c16.91-2.679 28.447-18.558 25.769-35.468-2.678-16.91-18.558-28.448-35.468-25.77Z"
        />
        <path
          fill="#fff"
          d="M13.81 34.594c.029 1.713.965 3.34 2.6 4.523 1.636 1.182 3.838 1.824 6.122 1.785 2.283-.04 4.461-.759 6.055-1.998 1.593-1.24 2.472-2.898 2.442-4.61M16.272 27.662c-.017-.951-.608-1.712-1.322-1.7-.713.013-1.278.794-1.261 1.745.017.95.608 1.712 1.322 1.7.713-.013 1.277-.794 1.26-1.745ZM30.048 27.421c-.017-.95-.609-1.712-1.322-1.7-.713.013-1.278.794-1.261 1.745.016.951.608 1.712 1.321 1.7.713-.013 1.278-.794 1.262-1.745Z"
        />
      </g>
    </g>
    <defs>
      <clipPath id="a">
        <path fill="#fff" d="M0 0h62v62H0z" />
      </clipPath>
    </defs>
  </svg>
)
export default AvatarPlaceholder
