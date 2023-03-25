//
//  ContentView.swift
//  SecretHitlerIOS
//
//  Created by Isaac Rae on 2/9/23.
//

import SwiftUI

struct ContentView: View {
    @StateObject var authenticator = Authenticator()
    let service = Service()
    var body: some View {
        VStack {
            if !authenticator.authenticated {
                AuthView(service, authenticator)
            } else {
                JoinView(service: service)
            }
        }
        .padding()
    }
}

struct ContentView_Previews: PreviewProvider {
    static var previews: some View {
        ContentView()
    }
}
