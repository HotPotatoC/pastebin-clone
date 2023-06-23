import AvatarPlaceholder from './avatar-placeholder'

const Header = () => {
  return (
    <header className="flex items-center justify-between py-6">
      <h1 className="font-fira-code text-xl font-semibold">Browse</h1>
      <AvatarPlaceholder />
    </header>
  )
}

export default Header
