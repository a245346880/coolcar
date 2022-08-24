package id

type AccountID string

func (a AccountID) String() string {
	return string(a)
}

type TripID string

func (t TripID) String() string {
	return string(t)
}

type IdentityID string

func (i IdentityID) String() string {
	return string(i)
}

type CarID string

func (i CarID) String() string {
	return string(i)
}

type BlobID string

func (i BlobID) String() string {
	return string(i)
}
