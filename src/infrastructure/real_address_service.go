package infrastructure

import (
	"fmt"
	"strconv"

	"github.com/mkrill/diPrototypeWire/src/infrastructure/client"
	"github.com/mkrill/diPrototypeWire/src/infrastructure/client/address_search"
	"github.com/mkrill/diPrototypeWire/src/models"
)

type RealAddressService struct {
	// addressLookupClient is only used internal of real_address_service
	addressLookupClient *client.AddressLookup
}

func ProvideRealAddressService(addressLookupClient *client.AddressLookup) *RealAddressService {
	return &RealAddressService{
		addressLookupClient: addressLookupClient,
	}
}

// naming conventions
// go-standard: r
// project-standard: ras (all caps of service name in small letters)
func (ras *RealAddressService) FullAddressInfo(street, streetNr, postCode, city string) (
	*models.Address,
	error,
) {
	// fill parameter structure for serve call parameters
	params := address_search.NewExtendedResolveFullAddressInfoUsingGETParams()
	params.SetStreetName(street)
	houseNumber, err := strconv.Atoi(streetNr)
	if err != nil {
		return nil, fmt.Errorf("error in converting parameter housenumber into int: %w", err)
	}
	params.SetHouseNumber(int32(houseNumber))
	params.SetZipCode(postCode)
	params.SetMunicipalityName(city)

	// call ExtendedResolveFullAddressInfoUsingGET service
	result, err := ras.addressLookupClient.AddressSearch.ExtendedResolveFullAddressInfoUsingGET(params)

	if err != nil {
		return nil, fmt.Errorf("error in calling realAddressService: %w", err)
	}

	return &models.Address{
		Street:   result.Payload.StreetName,
		StreetNr: strconv.Itoa(int(result.Payload.HouseNumber)),
		PostCode: result.Payload.ZipCode,
		City:     result.Payload.MunicipalityName,
		KlsId:    strconv.Itoa(int(result.Payload.KlsID)),
	}, nil
}
