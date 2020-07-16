package crew

import (
	pb "github.com/J-Thompson12/planet-express/ship/pkg/planetexpress"
)

func GetCrew() (*pb.GetCrewResponse, error) {
	crew := []*pb.Crew{}

	fryMap := make(map[string]string)
	fryMap["injured"] = "Dropped package on his foot"
	fry := pb.Crew{
		Member: "Fry",
		Status: fryMap,
	}

	benderMap := make(map[string]string)
	benderMap["Healthy"] = "Drank all the beer"
	bender := pb.Crew{
		Member: "Bender",
		Status: benderMap,
	}

	leelaMap := make(map[string]string)
	leelaMap["dead"] = "Killed by space bee"
	leela := pb.Crew{
		Member: "Leela",
		Status: leelaMap,
	}

	crew = append(crew, &fry)
	crew = append(crew, &bender)
	crew = append(crew, &leela)
	return &pb.GetCrewResponse{
		Crew: crew,
	}, nil
}
